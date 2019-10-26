# go-leaf-demo [![License](http://img.shields.io/badge/license-apache-blue.svg?style=flat-square)](https://github.com/chenwbyx/leafserver/blob/master/LICENSE)
 基于leaf搭建的游戏服框架

## Features
* 实现了登录流程
* 使用Protobuf数据格式传输
* conf配置文件：采用[go-ini/ini](https://github.com/go-ini/ini)
* log日志采用：采用[zap/zapcore](https://godoc.org/go.uber.org/zap/zapcore)和[natefinch/lumberjack.v2](https://gopkg.in/natefinch/lumberjack.v2)，参考[文章](https://studygolang.com/articles/14220)
* redis:采用[go-redis](https://github.com/go-redis/redis),对set和Hset进行封装，支持单机和集群
* Token方案:基于[jwt-go](https://github.com/dgrijalva/jwt-go)实现
* orm:[gorm](https://github.com/jinzhu/gorm)连MySQL
* 定时器：封装了[robfig/cron](https://github.com/robfig/cron)

### Requirments
Go version>=1.12 and GO111MODULE=on 