package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

const (
	redisKeyFormat = "%v:%v"
)

type Adapter interface {
	WithContext(ctx context.Context) Adapter
	Get(key string, v interface{}) error
	Set(key string, v interface{}, expiration time.Duration) error
}

type Redis struct {
	ctx    context.Context
	client *redis.Client
	lock   sync.Mutex
	prefix string
}

func NewRedisAdapter(client *redis.Client, prefix string) Adapter {
	return &Redis{
		client: client,
		prefix: prefix,
		lock:   sync.Mutex{},
	}
}

func (rc *Redis) WithContext(ctx context.Context) Adapter {
	rc.ctx = ctx
	return rc
}

func (rc *Redis) Get(key string, v interface{}) error {
	originalKey := fmt.Sprintf(redisKeyFormat, rc.prefix, key)
	rc.lock.Lock()
	defer rc.lock.Unlock()
	data, err := rc.client.Get(originalKey).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (rc *Redis) Set(key string, v interface{}, expiration time.Duration) error {
	originalKey := fmt.Sprintf(redisKeyFormat, rc.prefix, key)
	rc.lock.Lock()
	defer rc.lock.Unlock()
	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	_, err = rc.client.Set(originalKey, data, expiration).Result()
	return err
}
