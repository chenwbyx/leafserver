package internal

import (
	"fmt"
	"github.com/chenwbyx/leafserver/server/models"
	"github.com/chenwbyx/leafserver/server/msg"
	"github.com/chenwbyx/leafserver/server/pkg/logging"
	"github.com/chenwbyx/leafserver/server/pkg/util"
	"github.com/name5566/leaf/gate"
	"go.uber.org/zap"
	"reflect"
	"regexp"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.UserLogin{}, handleUserLogin)
	handler(&msg.UserRegister{}, handleUserRegister)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleUserRegister(args []interface{}) {
	m := args[0].(*msg.UserRegister)
	a := args[1].(gate.Agent)
	name := m.GetLoginName()
	pwd := m.GetLoginPW()
	logging.LoginLogger.Info("receive UserRegister", zap.String("LoginName",m.GetLoginName()))

	reg := regexp.MustCompile(`/^[a-zA-Z\d]\w{2,10}[a-zA-Z\d]$/`)
	matched := reg.FindString(name)
	if(matched!=" "){
		logging.LoginLogger.Debug("UserRegister name is licit")
	}
	ok, err := models.GetLoginByName(name)
	if err != nil {
		fmt.Println(err)
	}
	if ok == true {
		logging.LoginLogger.Debug("UserRegister find in fail" )
		retBuf := &msg.UserRegister_Result{
			DefaultReault: &msg.DefaultReault{
				RetResult: msg.Result_REGISTER_FAIL,
				ErrorInfo:"账号已存在！",
			},
		}
		a.WriteMsg(retBuf)
	}
	err = models.AddLogin(name,pwd)
	if err != nil {
		fmt.Println(err)
		logging.LoginLogger.Debug("UserRegister write in fail" )
		retBuf := &msg.UserRegister_Result{
			DefaultReault: &msg.DefaultReault{
				RetResult: msg.Result_REGISTER_FAIL,
				ErrorInfo:"注册失败，请稍后再试！",
			},
		}
		a.WriteMsg(retBuf)
	} else{
		logging.LoginLogger.Info("UserRegister write in success" )
		retBuf := &msg.UserRegister_Result{
			DefaultReault: &msg.DefaultReault{
				RetResult: msg.Result_REGISTER_SUCCESS,
			},
		}

		a.WriteMsg(retBuf)
	}
}

func handleUserLogin(args []interface{}) {
	m := args[0].(*msg.UserLogin)
	a := args[1].(gate.Agent)
	name := m.GetLoginName()
	pwd := m.GetLoginPW()
	logging.LoginLogger.Info("receive UserRegister", zap.String("LoginName", m.GetLoginName()))

	ok, err := models.CheckLogin(name, pwd)
	if err != nil {
		fmt.Println(err)
	}
	if ok == true {
		logging.LoginLogger.Info("UserLogin find in success" )
		token, err := util.GenerateToken(name, pwd)
		if err != nil {
			//return
		}
		retBuf := &msg.UserLogin_Result{
			DefaultReault: &msg.DefaultReault{
				RetResult: msg.Result_LOGIN_SUCCESS,
				ErrorInfo:"登录成功！",
			},
			Token: token,
		}
		a.WriteMsg(retBuf)
	}

	retBuf := &msg.UserLogin_Result{
		DefaultReault: &msg.DefaultReault{
			RetResult: msg.Result_LOGIN_SUCCESS,
			ErrorInfo:"登陆失败，请稍后再试！",
		},
	}
	a.WriteMsg(retBuf)
}