package inits

import (
	"context"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/common/global"
	"github.com/go-redis/redis/v9"
	"time"
)

func Redis(conn string) error {
	// Parse connect string
	redisConfig, err := redis.ParseURL(conn)
	if err != nil {
		return fmt.Errorf("failed to parse redis connection string: %v", err)
	}

	// Connect to server
	global.Redis = redis.NewClient(redisConfig)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Try connection
	err = global.Redis.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %v", err)
	}

	return nil
}
