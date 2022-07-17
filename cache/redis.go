package cache

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

// TODO redis cache

type Redis struct {
	rdb *redis.Client
}

// NewRedis 实例化
func NewRedis(rdb *redis.Client) *Redis {
	return &Redis{rdb}
}

// Get 获取一个值
func (r *Redis) Get(key string) (val interface{}) {
	val, err := r.rdb.Get(key).Result()
	if err != redis.Nil {
		return val
	}
	return nil
}

// Set 设置一个值
func (r *Redis) Set(key string, val interface{}, timeout time.Duration) (err error) {
	return r.rdb.Set(key, val, timeout).Err()
}

// IsExist 判断key是否存在
func (r *Redis) IsExist(key string) bool {
	i, err := r.rdb.Exists(key).Result()
	log.Println(i, err)
	return i > 0
}

// Delete 删除
func (r *Redis) Delete(key string) error {
	return nil
}
