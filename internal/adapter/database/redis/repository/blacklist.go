package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type BlacklistRepository struct {
	redis *redis.Client
}

func NewBlacklistRepository(redis *redis.Client) *BlacklistRepository {
	return &BlacklistRepository{redis: redis}
}

func (r *BlacklistRepository) AddToBlackList(token string, expire time.Time) error {
	ctx := context.Background()
	key := "bl_" + token

	if err := r.redis.Set(ctx, key, token, 0).Err(); err != nil {
		fmt.Println("SET : ", err)
		return err
	}

	if err := r.redis.ExpireAt(ctx, key, expire).Err(); err != nil {
		fmt.Println("ExpireAt : ", err)
		return err
	}
	return nil
}

func (r *BlacklistRepository) IsInBlackList(token string) bool {
	ctx := context.Background()
	key := "bl_" + token
	return r.redis.Exists(ctx, key).Val() == 1
}
