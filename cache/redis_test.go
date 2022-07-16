package cache

import (
	"log"
	"testing"
	"time"

	"github.com/bugstark/at/datas"
)

func TestRedis_Set(t *testing.T) {
	datas.InitRedis("127.0.0.1", "", 6379, 0)
	ca := NewRedis(datas.Redis)
	log.Println(ca.Set("test", "123123123123123", time.Second*10))
	log.Println(ca.IsExist("tes2t"))
	log.Println(ca.Get("test"))
	log.Println(ca.Delete("test"))
}
