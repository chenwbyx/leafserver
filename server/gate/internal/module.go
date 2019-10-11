package internal

import (
	"github.com/chenwbyx/leafserver/server/game"
	"github.com/chenwbyx/leafserver/server/msg"
	"github.com/chenwbyx/leafserver/server/pkg/setting"
	"github.com/name5566/leaf/gate"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      setting.ServerSetting.MaxConnNum,
		PendingWriteNum: setting.PendingWriteNum,
		MaxMsgLen:       setting.MaxMsgLen,
		WSAddr:          setting.ServerSetting.WSAddr,
		HTTPTimeout:     setting.HTTPTimeout,
		CertFile:        setting.ServerSetting.CertFile,
		KeyFile:         setting.ServerSetting.KeyFile,
		TCPAddr:         setting.ServerSetting.TCPAddr,
		LenMsgLen:       setting.LenMsgLen,
		LittleEndian:    setting.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}
}
