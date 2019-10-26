package internal

import (
	"github.com/chenwbyx/leafserver/server/models"
	"github.com/chenwbyx/leafserver/server/msg"
	"github.com/chenwbyx/leafserver/server/pkg/e"
	"github.com/chenwbyx/leafserver/server/pkg/logging"
	"github.com/chenwbyx/leafserver/server/pkg/util"
	"github.com/name5566/leaf/gate"
	"go.uber.org/zap"
	"log"
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
		log.Println(err)
	}
	if ok == true {
		logging.LoginLogger.Debug("UserRegister find in fail" )
		retBuf := &msg.UserRegister_Result{
			DefaultReault: &msg.DefaultReault{
				RetResult: e.ACCOUNT_EXIST,
				ErrorInfo: e.GetMsg(e.ACCOUNT_EXIST),
			},
		}
		a.WriteMsg(retBuf)
	}
	err = models.AddLogin(name,pwd)
	if err != nil {
		log.Println(err)
		logging.LoginLogger.Debug("UserRegister write in fail" )
		retBuf := &msg.UserRegister_Result{
			DefaultReault: &msg.DefaultReault{
				RetResult: e.REGISTRE_FAIL,
				ErrorInfo: e.GetMsg(e.REGISTRE_FAIL),
			},
		}
		a.WriteMsg(retBuf)
	} else{
		logging.LoginLogger.Info("UserRegister write in success" )
		retBuf := &msg.UserRegister_Result{
			DefaultReault: &msg.DefaultReault{
				RetResult: e.REGISTRE_SUCCESS,
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
		log.Println(err)
	}
	if ok == true {
		logging.LoginLogger.Info("UserLogin find in success" )
		token, err := util.GenerateToken(name, pwd)
		if err != nil {
			//return
		}
		retBuf := &msg.UserLogin_Result{
			DefaultReault: &msg.DefaultReault{
				RetResult: e.LOGIN_SUCCESS,
				ErrorInfo: e.GetMsg(e.LOGIN_SUCCESS),
			},
			Token: token,
		}
		a.WriteMsg(retBuf)
	}

	retBuf := &msg.UserLogin_Result{
		DefaultReault: &msg.DefaultReault{
			RetResult: e.LOGIN_FAIL,
			ErrorInfo: e.GetMsg(e.LOGIN_FAIL),
		},
	}
	a.WriteMsg(retBuf)
}