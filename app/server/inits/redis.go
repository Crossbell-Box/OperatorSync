package inits

import (
	"context"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/go-redis/redis/v9"
	"time"
)

func Redis() error {
	// Parse connect string
	redisConfig, err := redis.ParseURL(config.Config.RedisConnString)
	if err != nil {
		return fmt.Errorf("failed to parse redis connection string: %w", err)
	}

	// Connect to server
	global.Redis = redis.NewClient(redisConfig)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Try connection
	err = global.Redis.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	return nil
}
