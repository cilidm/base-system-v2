package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	DBConn    *gorm.DB
	ZapLog    *zap.Logger
	RedisConn *redis.Client
)
