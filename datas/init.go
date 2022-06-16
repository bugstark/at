package datas

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Redis *redis.Client

func InitDB(drive, dsn string, debug bool) {
	if drive == "mysql" {
		DB = initmysql(dsn, debug)
		return
	}
	DB = initsqlite(debug)
}

func InitRedis(host, pass string, port, db int) {
	Redis = initredis(host, pass, port, db)
}
