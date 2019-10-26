package cron

import (
	"github.com/chenwbyx/leafserver/server/pkg/gredis"
	"github.com/chenwbyx/leafserver/server/pkg/logging"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"runtime/debug"
)

// 秒(Seconds)		0-59
// 分(Minutes)		0-59
// 时(Hours)		0-23
// 日(Day of month)		1-31
// 月(Month)		1-12 or JAN-DEC
// 星期(Day of week) 0-6
//
var crontabGame *cron.Cron

func Setup() error {
	var jobs = []struct {
		spec   string
		worker func() error
		name   string
	}{
		// 注册定时任务
		{"*/5 * * * * ?", jobUpdateRedisServerInfo, "UpdateRedisServerInfo"},
	}

	crontabGame = cron.New()
	for _, job := range jobs {
		err := crontabGame.AddFunc(job.spec, cronjobWrapper(job.name, job.worker))
		if err != nil {
			return err
		}
	}
	crontabGame.Start()
	// 启动时需要先执行的job
	// go jobCleanMails()
	return nil
}

func cronjobWrapper(name string, worker func() error) func() {
	return func() {
		defer func() {
			if err := recover(); err != nil {
				logging.AppLogger.Error("cronjob", zap.String("cronjob",name), zap.ByteString("stack", debug.Stack()))
			}
		}()
		err := worker()
		if err != nil {
			logging.AppLogger.Error("cronjob ", zap.String("cronjob",name))
		}
	}
}

//定时ping() Redis
func jobUpdateRedisServerInfo() error {

	{
		result := gredis.Global_redis.Ping()
		if result.Err() != nil {
			return result.Err()
		}
	}

	return nil
}