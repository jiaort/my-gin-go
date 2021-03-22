package initialize

import (
	"my-gin-go/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.G_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.G_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.G_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.G_REDIS = client
	}
}
