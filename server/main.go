package main

import (
	"github.com/chenwbyx/leafserver/server/game"
	"github.com/chenwbyx/leafserver/server/gate"
	"github.com/chenwbyx/leafserver/server/login"
	"github.com/chenwbyx/leafserver/server/pkg/setting"
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = setting.ServerSetting.LogLevel
	lconf.LogPath = setting.AppSetting.RuntimeRootPath + setting.AppSetting.LogSavePath
	lconf.LogFlag = setting.AppSetting.LogFlag
	lconf.ConsolePort = setting.ServerSetting.ConsolePort
	lconf.ProfilePath = setting.ServerSetting.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}