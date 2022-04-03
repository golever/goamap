package excel

type Poi struct {
	Seq       string
	NType     string
	BCategory string
	MCategory string
	SCategory string
}

var poiCodeFilePath = "excel_file/poi_code.xlsx"

func ReadPoiExcel() []Poi {
	var p []Poi
	rows := readerexcel(poiCodeFilePath, "amap_poicode")
	for _, row := range rows {
		po := Poi{
			Seq:       row[0],
			NType:     row[1],
			BCategory: row[2],
			MCategory: row[3],
			SCategory: row[4],
		}
		p = append(p, po)
	}
	return p
}
