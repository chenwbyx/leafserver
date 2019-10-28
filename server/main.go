package main

import (
	"github.com/chenwbyx/leafserver/server/cron"
	"github.com/chenwbyx/leafserver/server/game"
	"github.com/chenwbyx/leafserver/server/gate"
	"github.com/chenwbyx/leafserver/server/login"
	"github.com/chenwbyx/leafserver/server/models"
	"github.com/chenwbyx/leafserver/server/pkg/gredis"
	"github.com/chenwbyx/leafserver/server/pkg/logging"
	"github.com/chenwbyx/leafserver/server/pkg/setting"
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func init()  {
	setting.Setup()
	logging.Setup()
	models.Setup()
	gredis.Setup()
	cron.Setup()
}

func main() {
	lconf.LogLevel = setting.ServerSetting.LogLevel
	lconf.LogPath = ""
	lconf.LogFlag = setting.AppSetting.LogFlag
	lconf.ConsolePort = setting.ServerSetting.ConsolePort
	lconf.ProfilePath = setting.ServerSetting.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
