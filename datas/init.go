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

// Query 查询参数
type Query struct {
	Fields string `gorm:"-" form:"fields" json:"-"`
	Sort   string `gorm:"-" form:"sort" json:"-"`
	Order  string `gorm:"-" form:"order" json:"-"`
	Size   int    `gorm:"-" form:"size" json:"-"`
	Page   int    `gorm:"-" form:"page" json:"-"`
}

type Model struct {
	ID        string         `gorm:"PRIMARY_KEY;UNIQUE;type:varchar(20);column:id;comment:'唯一ID'" json:"id,omitempty" form:"id"`
	CreatedAt int64          `gorm:"type:int(10);column:createdat;comment:'创建时间'" json:"createdat,omitempty" form:"createdat"`
	UpdatedAt int64          `gorm:"type:int(10);column:updatedat;comment:'更新时间'" json:"updatedat,omitempty" form:"updatedat"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:'删除时间';column:deletedat" json:"deletedat,omitempty" form:"deletedat"`
}
