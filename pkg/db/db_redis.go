package db

import (
	"github.com/go-redis/redis"
	"github.com/sharch/sim/config"
	"github.com/sharch/sim/pkg/logger"
	"github.com/sharch/sim/pkg/util"
)

var (
	RedisCli  *redis.Client
	RedisUtil *util.RedisUtil
)

func init() {
	InitRedis(config.Config.RedisHost, config.Config.RedisPassword)
}

// InitRedis 初始化Redis
func InitRedis(addr, password string) {
	logger.Logger.Info("init redis")
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       0,
		Password: password,
	})

	_, err := RedisCli.Ping().Result()
	if err != nil {
		panic(err)
	}

	RedisUtil = util.NewRedisUtil(RedisCli)
	logger.Logger.Info("init redis ok")
}
