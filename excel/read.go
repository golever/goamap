package excel

import (
	"github.com/xuri/excelize/v2"
	"log"
)

func readerexcel(filepath, sheet string) [][]string {
	//flg := flag.String("city", "excel_file/city_code.xlsx", "city ad code and city code")
	//flag.Parse()
	file, err := excelize.OpenFile(filepath)
	if err != nil {
		log.Printf("read city_code.xlsx err %s", err)
	}
	defer func(file *excelize.File) {
		err := file.Close()
		if err != nil {
			log.Printf("file close err %s", err)
		}
	}(file)
	rows, _ := file.GetRows(sheet)
	return rows
}
