package req

import (
	"encoding/json"
	"goamap.mod/conf"
	"goamap.mod/excel"
	"goamap.mod/repo"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var data repo.Response

type PoiParam struct {
	offset   int
	page     int
	keywords string
	type_    string
	city     string
}

/**
* 因poi服务类型较多，这里只获取类型为学校的数据
* 获取高德地图所有类型数据，可调用ReadPoiExcel()函数来获取POI类型code
 */
var poiCode = [8]string{"141200", "141201", "141202", "141203", "141204", "141205", "141206", "141207"}

func DoGetPoi(city string) {

	cits := excel.ReadCityExcel()
	var cs string
	for _, c := range cits {
		if strings.EqualFold(city, c.Name) {
			cs = c.CityCode
		}
	}

	k := conf.Cfg.Guide.Key

	for _, pc := range poiCode {
		count := 1
		m := map[string]string{
			"key":    k,
			"types":  pc,
			"city":   cs,
			"offset": "10",
			"page":   strconv.Itoa(count),
		}
		doGet(&m)
		for {
			dc, _ := strconv.Atoi(data.Count)
			if count < dc/10 {
				doGet(&m)
				count++
				continue
			}
			break
		}

	}

}

func doGet(m *map[string]string) {

	realpath := urlEncode(m)
	log.Printf("url:%s", realpath)
	resp, err := http.Get(realpath)
	if err != nil {
		log.Printf("the request failed url:%s", realpath)
		return
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	er := json.Unmarshal(content, &data)
	if er != nil {
		log.Printf("data conversion exception err %s", err)
		return
	}
	if data.Info != "ok" {
		log.Printf("result data err %s", data.Info)
		return
	}
	repo.SavePoi(&data)
}

func urlEncode(m *map[string]string) string {
	par, _ := url.Parse(conf.Cfg.Guide.Url)
	p := url.Values{}

	for k, v := range *m {
		p.Set(k, v)
	}

	par.RawPath = p.Encode()
	return par.String()
}
