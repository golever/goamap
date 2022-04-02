package conf

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Mongo struct {
		Url        string `yaml:"url"`
		Db         string `yaml:"db"`
		Collection string `yaml:"collection"`
	} `yaml:"mongo"`

	Guide struct {
		Url string `yaml:"url"`
		Key string `yaml:"key"`
	} `yaml:"guide"`
}

var Cfg Config

func init() {
	Configures()
}

func Configures() {
	flg := flag.String("config", "yaml/config.yml", "file config")
	flag.Parse()
	b, err := ioutil.ReadFile(*flg)
	if err != nil {
		log.Fatalf("read file config.yml error %s/n ", err)
	}
	er := yaml.Unmarshal(b, &Cfg)
	if er != nil {
		log.Fatalf("failed to parse config. cause by %s", err.Error())
	}
	log.Printf("config file content: %v", Cfg)
}
