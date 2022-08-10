package excel

import (
	"log"
	"testing"

	_ "github.com/pbnjay/grate/xls"
	_ "github.com/pbnjay/grate/xlsx"
)

func TestRead(t *testing.T) {
	log.Println(ReadExcel2("检测数据导出2.xlsx", 0))
}
