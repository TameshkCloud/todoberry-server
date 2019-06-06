package redis_client

import (
	"time"

	"github.com/TameshkCloud/todoberry-server/config"
	"github.com/go-redis/redis"
)

const DEFAULT_REDIS string = "redis_default"

var client redis.UniversalClient

func New(key string) redis.UniversalClient {
	dsn := config.Redis.DSN
	switch key {
	case DEFAULT_REDIS:
		dsn = config.Redis.DSN
	}
	client = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        dsn,
		PoolTimeout:  time.Duration(config.Redis.Timeouts.InternalPoolTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.Redis.Timeouts.IdleTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Redis.Timeouts.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Redis.Timeouts.WriteTimeout) * time.Second,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}