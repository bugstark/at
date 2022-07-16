package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	rdb *redis.Client
}

// NewRedis 实例化
func NewRedis(rdb *redis.Client) *Redis {
	return &Redis{rdb}
}

// Get 获取一个值
func (r *Redis) Get(key string) (val interface{}) {
	return
}

// Set 设置一个值
func (r *Redis) Set(key string, val interface{}, timeout time.Duration) (err error) {
	return
}

// IsExist 判断key是否存在
func (r *Redis) IsExist(key string) bool {
	return false
}

// Delete 删除
func (r *Redis) Delete(key string) error {
	return nil
}
