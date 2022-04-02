package excel

type Poi struct {
	Seq       string
	NType     string
	BCategory string
	MCategory string
	SCategory string
}

var PoiCode []Poi
var poiCodeFilePath = "excel_file/poi_code.xlsx"

func init() {
	readPoiExcel()
}

func readPoiExcel() {
	rows := readerexcel(poiCodeFilePath, "amap_poicode")
	for _, row := range rows {
		po := Poi{
			Seq:       row[0],
			NType:     row[1],
			BCategory: row[2],
			MCategory: row[3],
			SCategory: row[4],
		}
		PoiCode = append(PoiCode, po)
	}
}
