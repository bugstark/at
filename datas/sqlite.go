package datas

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initsqlite(debug bool) *gorm.DB {
	var level = logger.Silent
	if debug {
		level = logger.Info
	}
	initdb, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold:             5 * time.Second, // 慢 SQL 阈值
				LogLevel:                  level,           // 日志级别
				IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,            // 禁用彩色打印
			}),
	})
	if err != nil {
		log.Panicln(err)
		return nil
	}
	return initdb
}
