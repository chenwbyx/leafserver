package gredis

import (
	"github.com/chenwbyx/leafserver/server/pkg/logging"
	"github.com/chenwbyx/leafserver/server/pkg/setting"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"time"
)

type RedisDb struct {
}


var Global_redis = &RedisDb{}

//redis单机客户端
var client *redis.Client

//redis集群客户端
var redisClusterClient *redis.ClusterClient



//redis设置
func Setup() bool {

	//判断是否为集群配置
	if setting.RedisSetting.Cluster {
		//ClusterClient是一个Redis集群客户机，表示一个由0个或多个底层连接组成的池。它对于多个goroutine的并发使用是安全的。
		redisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
			Password: setting.RedisSetting.Password,
			Addrs:    setting.RedisSetting.Hosts,
		})
		//Ping
		ping, err := redisClusterClient.Ping().Result()
		if err != nil {
			logging.AppLogger.Error("Redis Ping", zap.String("Addr", setting.RedisSetting.Host), zap.Error(err))
			return false
		}
		logging.AppLogger.Info("redis ping result", zap.String("Addr", setting.RedisSetting.Host), zap.String("result", ping))
		return true

	} else {
		//Redis客户端，由零个或多个基础连接组成的池。它对于多个goroutine的并发使用是安全的。
		//更多参数参考Options结构体
		client = redis.NewClient(&redis.Options{
			Addr:     setting.RedisSetting.Host,
			Password: setting.RedisSetting.Password, // no password set
			DB:       setting.RedisSetting.DB,       // use default DB
		})
		//Ping
		ping, err := client.Ping().Result()
		if err != nil {
			logging.AppLogger.Error("Redis Ping", zap.String("Addr", setting.RedisSetting.Host), zap.Error(err))
			return false
		}
		logging.AppLogger.Info("redis ping result", zap.String("Addr", setting.RedisSetting.Host), zap.String("result", ping))
		return true
	}

}

//set
func  (this *RedisDb)  Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if setting.RedisSetting.Cluster {
		return redisClusterClient.Set(key, value, expiration)
	} else {
		return client.Set(key, value, expiration)
	}
}

//get
func  (this *RedisDb)  Get(key string) *redis.StringCmd {
	if setting.RedisSetting.Cluster {
		return redisClusterClient.Get(key)
	} else {
		return client.Get(key)
	}
}

// Ping 检测链接是否正常
// 一段时间后redis操作前，需要调用一下Ping(), 如果已经发生断开，会自动先触发重连，避免后继操作出错
func  (this *RedisDb)  Ping() *redis.StatusCmd {
	var result *redis.StatusCmd
	if setting.RedisSetting.Cluster {
		result = redisClusterClient.Ping()		//第一次出错，尝试再连接一次，确保闪断情况继续连上
		if result.Err() != nil {
			result = redisClusterClient.Ping()
		}
	} else {
		result = client.Ping()		//第一次出错，尝试再连接一次，确保闪断情况继续连上
		if result.Err() != nil {
			result = client.Ping()
		}
	}

	return result
}

//HSet 某个的hash表的字段设置
func  (this *RedisDb)  HSet(key string, field string, value interface{}) *redis.BoolCmd {
	if setting.RedisSetting.Cluster {
		return redisClusterClient.HSet(key, field, value)
	} else {
		return client.HSet(key, field, value)
	}
}

//HGet 某个的hash表的字段获取
func  (this *RedisDb)  HGet(key string, field string) *redis.StringCmd {
	if setting.RedisSetting.Cluster {
		return redisClusterClient.HGet(key, field)
	} else {
		return client.HGet(key, field)
	}
}

//HDel 某个的hash表的删除某些字段
func  (this *RedisDb)  HDel(key string, fields ...string) *redis.IntCmd {
	if setting.RedisSetting.Cluster {
		return redisClusterClient.HDel(key, fields ...)
	} else {
		return client.HDel(key, fields ...)
	}
}

//订阅
func (this *RedisDb) Subscribe(channels ...string) *redis.PubSub {
	if setting.RedisSetting.Cluster {
		return redisClusterClient.Subscribe(channels...)
	} else {
		return client.Subscribe(channels...)
	}
}

//发布
func  (this *RedisDb) Publish(channel string, message interface{}) *redis.IntCmd {
	if setting.RedisSetting.Cluster {
		return redisClusterClient.Publish(channel, message)
	} else {
		return client.Publish(channel, message)
	}
}

//Close
func (this *RedisDb) Close() error{
	if setting.RedisSetting.Cluster {
		return redisClusterClient.Close();
	} else {
		return client.Close();
	}
}
