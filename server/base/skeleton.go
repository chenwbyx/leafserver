package base

import (
	"github.com/chenwbyx/leafserver/server/pkg/setting"
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/module"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              setting.GoLen,
		TimerDispatcherLen: setting.TimerDispatcherLen,
		AsynCallLen:        setting.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(setting.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}
