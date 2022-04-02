package req

import (
	"encoding/json"
	"fmt"
	"goamap.mod/conf"
	"goamap.mod/repo"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var data repo.Response

func Do(keywords string, tp string, city string) {

	realpath := urlEncode(keywords, tp, city)
	resp, err := http.Get(realpath)
	if err != nil {
		log.Fatalf("the request failed url:%s", realpath)
		return
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	er := json.Unmarshal(content, &data)
	if er != nil {
		log.Fatalf("data conversion exception err %s", err)
		return
	}
	fmt.Println(data)
	if data.Info != "ok" {
		log.Fatalf("result data err %s", data.Info)
		return
	}
	repo.SavePoi(&data)
}

func urlEncode(keywords string, tp string, city string) string {
	par, _ := url.Parse(conf.Cfg.Guide.Url)
	p := url.Values{}
	p.Set("key", conf.Cfg.Guide.Key)
	p.Set("keywords", keywords)
	p.Set("types", tp)
	p.Set("city", city)

	par.RawPath = p.Encode()
	return par.String()
}
