package gredis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	RDB *redis.Client // Redis连接池
	ctx = context.Background()
)
