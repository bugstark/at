package datas

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func initredis(host, pass string, port, db int) *redis.Client {
	// 设置redis客户端
	log.Printf("连接redis [%s:%d]", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: pass, // no password set
		DB:       db,   // use default DB
	})
	// 测试连接
	_, e := rdb.Ping().Result()
	if e != nil {
		log.Fatalf("连接redis失败! [%s:%d]", host, port)
	}
	return rdb
}
