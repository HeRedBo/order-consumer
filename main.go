package main

import (
	"github.com/HeRedBo/pkg/cache"
	"github.com/HeRedBo/pkg/es"
	"github.com/HeRedBo/pkg/logger"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"order-consumer/global"
)

func init() {
	global.LoadConfig()
	global.LOG = global.SetupLogger()
	initRedisClient()
	initESClient()

}

func initRedisClient() {
	redisCfg := global.CONFIG.Redis
	opt := redis.Options{
		Addr:        redisCfg.Host,
		Password:    redisCfg.Password,
		IdleTimeout: redisCfg.IdleTimeout,
	}
	//redisTrace := trace.Cache{
	//	Name:                  "redis",
	//	SlowLoggerMillisecond: 500,
	//	Logger:                logger.GetLogger(),
	//	AlwaysTrace:           global.CONFIG.App.RunMode == conf.RunModeDev,
	//}
	err := cache.InitRedis(cache.DefaultRedisClient, &opt)
	if err != nil {
		logger.Error("redis init error", zap.Error(err))
		panic("initRedisClient error")
	}
}

func initESClient() {
	err := es.InitClientWithOptions(es.DefaultClient, global.CONFIG.Elasticsearch.Hosts,
		global.CONFIG.Elasticsearch.Username, global.CONFIG.Elasticsearch.Password,
		es.WithScheme("https"))
	if err != nil {
		panic(err)
	}
}

func main() {

}
