module github.com/chenwbyx/leafserver

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/go-ini/ini v1.48.0 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/name5566/leaf v0.0.0-20181103040206-1364c176dfbd
	github.com/pkg/errors v0.8.1 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/unknwon/com v1.0.1 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.2.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	gopkg.in/ini.v1 v1.48.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace (
	github.com/chenwbyx/leafserver/server/base => ./server/base
	github.com/chenwbyx/leafserver/server/conf => ./server/conf
	github.com/chenwbyx/leafserver/server/game => ./server/game
	github.com/chenwbyx/leafserver/server/gate => ./server/gate
	github.com/chenwbyx/leafserver/server/login => ./server/login
	github.com/chenwbyx/leafserver/server/msg => ./server/msg
	github.com/chenwbyx/leafserver/server/pkg => ./server/pkg
)
