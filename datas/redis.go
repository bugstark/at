package datas

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func Initredis(host, pass string, port, db int) {
	log.Printf("连接redis [%s:%d]", host, port)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: pass, // no password set
		DB:       db,   // use default DB
	})
	// 测试连接
	_, e := Rdb.Ping().Result()
	if e != nil {
		log.Fatalf("连接redis失败! [%s:%d]", host, port)
	}
}
