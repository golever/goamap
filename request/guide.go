package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const u = "https://restapi.amap.com/v3/place/text?parameters"
const key = "264b999f57395fbbee6bcbe002115b7a"

type Params struct {
	Keywords string
	Types    string
}

func Send(p Params) string {

	params := url.Values{}
	Url, err := url.Parse(u)
	if err != nil {
		return err.Error()
	}
	params.Set("key", key)
	params.Set("keywords", p.Keywords)
	params.Set("types", p.Types)

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println("url:{}", urlPath)

	resp, err := http.Get(urlPath)

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
