package excel

import (
	"github.com/pbnjay/grate"
	_ "github.com/pbnjay/grate/xls"
	_ "github.com/pbnjay/grate/xlsx"
)

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
