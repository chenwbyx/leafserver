module github.com/chenwbyx/leafserver

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-ini/ini v1.48.0
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/name5566/leaf v0.0.0-20181103040206-1364c176dfbd
	github.com/pkg/errors v0.8.1 // indirect
	github.com/robfig/cron v1.2.0
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/unknwon/com v1.0.1 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.2.0 // indirect
	go.uber.org/zap v1.10.0
	gopkg.in/ini.v1 v1.48.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/chenwbyx/leafserver/server/base => ./server/base
	github.com/chenwbyx/leafserver/server/conf => ./server/conf
	github.com/chenwbyx/leafserver/server/game => ./server/game
	github.com/chenwbyx/leafserver/server/gate => ./server/gate
	github.com/chenwbyx/leafserver/server/login => ./server/login
	github.com/chenwbyx/leafserver/server/models => ./server/models
	github.com/chenwbyx/leafserver/server/msg => ./server/msg
	github.com/chenwbyx/leafserver/server/pkg => ./server/pkg
)
