package request

import (
	"fmt"
	"goamap.mod/conf"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Params struct {
	Keywords string
	Types    string
}

func Send(p Params) string {

	pt := urlParamEncode(p)
	resp, err := http.Get(pt)

	defer resp.Body.Close()

	if err != nil {
		fmt.Println("error info {}", err)
		return err.Error()
	}

	return result(resp.Body)
}

func result(b io.ReadCloser) string {
	body, _ := ioutil.ReadAll(b)
	return string(body)
}

func urlParamEncode(p Params) string {
	var c conf.GuideConf
	c.Build()
	fmt.Println("url:" + c.Url)

	params := url.Values{}
	Url, err := url.Parse(c.Url)
	if err != nil {
		log.Println("param error,info {}", err)
	}

	params.Set("key", c.Key)
	params.Set("keywords", p.Keywords)
	params.Set("types", p.Types)

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	return urlPath
}
