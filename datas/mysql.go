package datas

import (
	"log"
	"os"
	"time"

	"github.com/rs/xid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func initmysql(dsn string, debug bool) *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   false, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true,  // 根据版本自动配置
	}
	var level = logger.Silent
	if debug {
		level = logger.Info
	}
	initdb, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
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
		log.Panicln("初始化数据库失败:" + err.Error())
		return nil
	}
	under, err := initdb.DB()
	if err != nil {
		log.Panicln("初始化数据库失败:" + err.Error())
	}
	under.SetConnMaxLifetime(time.Hour * 7)
	under.SetMaxIdleConns(2)
	under.SetMaxOpenConns(30)
	err = initdb.Callback().Create().Before("gorm:create").Register("ID", func(db *gorm.DB) {
		if db.Statement.Table == "Sys_Auth" {
			return
		}
		db.Statement.SetColumn("ID", xid.New().String())
	})
	if err != nil {
		log.Fatal("creat callback err", err)
	}
	log.Println("数据库初始化成功")
	return initdb

}
