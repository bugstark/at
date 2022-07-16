package cache

import (
	"encoding/json"
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
func (r *Redis) Get(key string) interface{} {
	var data []byte
	var err error
	if data, err = r.rdb.Get(key).Bytes(); err != nil {
		return nil
	}
	var reply interface{}
	if err = json.Unmarshal(data, &reply); err != nil {
		return nil
	}
	return reply
}

// Set 设置一个值
func (r *Redis) Set(key string, val interface{}, timeout time.Duration) (err error) {
	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return
	}
	return r.rdb.SetXX(key, data, timeout).Err()
}

// IsExist 判断key是否存在
func (r *Redis) IsExist(key string) bool {
	i, _ := r.rdb.Exists(key).Result()
	return i > 0
}

// Delete 删除
func (r *Redis) Delete(key string) error {
	return r.rdb.Del(key).Err()
}
