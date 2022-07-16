package datas

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func initredis(host, pass string, port, db int) *redis.Client {
	log.Printf("连接redis [%s:%d]", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Password:     pass, // no password set
		DB:           db,   // use default DB
		DialTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Second,
		PoolSize:     50,
		MaxConnAge:   10 * time.Second,
		IdleTimeout:  8 * time.Second,
	})
	// 测试连接
	_, e := rdb.Ping().Result()
	if e != nil {
		log.Panicf("连接redis失败! [%s:%d]", host, port)
	}
	return rdb
}
