package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb        *redis.Client
	defaultTTL time.Duration
}

func New(rdb *redis.Client, defaultTTL time.Duration) *Cache {
	return &Cache{
		rdb:        rdb,
		defaultTTL: defaultTTL,
	}
}

func Key(parts ...any) string {
	strs := make([]string, len(parts))

	for _, part := range parts {
		strs = append(strs, fmt.Sprint(part))
	}

	return strings.Join(strs, ":")
}

func (c *Cache) Get(ctx context.Context, key string, out any) (bool, error) {
	raw, err := c.rdb.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal(raw, out); err != nil {
		return false, err
	}

	return true, nil
}

func (c *Cache) Set(ctx context.Context, key string, value any) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, key, raw, c.defaultTTL).Err()
}

func (c *Cache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	return c.rdb.Del(ctx, keys...).Err()
}
