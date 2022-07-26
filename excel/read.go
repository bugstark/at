package excel

import (
	"log"

	"github.com/pbnjay/grate"
	_ "github.com/pbnjay/grate/xls"
	_ "github.com/pbnjay/grate/xlsx"
	"github.com/xuri/excelize/v2"
)

// 在win下使用疑似有问题，暂未排查
func Read(filepath string, sheets_index int) (res [][]string, err error) {
	wb, err := grate.Open(filepath)
	if err != nil {
		return nil, err
	}
	sheets, err := wb.List()
	if err != nil {
		return nil, err
	}
	sheet, err := wb.Get(sheets[sheets_index])
	if err != nil {
		return nil, err
	}
	var temp = [][]string{}
	for sheet.Next() {
		row := sheet.Strings()
		temp = append(temp, row)
	}
	wb.Close()
	return temp, nil
}

func ReadExcel2(filepath string, sheets int) (res [][]string, err error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()
	res, err = f.GetRows(f.GetSheetName(sheets))
	return
}
