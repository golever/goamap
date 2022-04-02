package excel

type CityCode struct {
	Name     string
	AdCode   string
	CityCode string
}

var City []CityCode
var cityCodeFilePath = "excel_file/city_code.xlsx"

func init() {
	readCityExcel()
}

func readCityExcel() {
	rows := readerexcel(cityCodeFilePath, "Sheet1")
	for _, row := range rows {
		if len(row) < 3 {
			row = append(row, "0")
		}
		cc := CityCode{
			Name:     row[0],
			AdCode:   row[1],
			CityCode: row[2],
		}
		City = append(City, cc)
	}

}
