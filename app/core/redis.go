package core

import (
	"github.com/cilidm/base-system-v2/app/global"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     Conf.Redis.RedisAddr,
		Password: Conf.Redis.RedisPWD, // no password set
		DB:       Conf.Redis.RedisDB,  // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.ZapLog.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.ZapLog.Info("redis connect ping response:", zap.String("pong", pong))
		global.RedisConn = client
	}
}
