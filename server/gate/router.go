package gate

import (
	"github.com/chenwbyx/leafserver/server/login"
	"github.com/chenwbyx/leafserver/server/msg"
)

func init() {
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.UserLogin{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.UserRegister{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.UserST{}, login.ChanRPC)
}
