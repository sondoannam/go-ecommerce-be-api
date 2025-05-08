package initialize

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sondoannam/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Db,      // use default DB
		PoolSize: r.PoolSize, // set max number of idle connections
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initializtion error: ", zap.Error(err))
	}

	fmt.Println("Redis connection established successfully!")
	global.Rdb = rdb

	redisTest()
}

func redisTest() {
	// Set a key-value pair in Redis
	err := global.Rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println("Error setting key:", zap.Error(err))
		return
	}

	// Get the value back from Redis
	value, err := global.Rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("Error getting key:", zap.Error(err))
		return
	}

	global.Logger.Info("Value from Redis:", zap.String("key", "key"), zap.String("value", value))
}