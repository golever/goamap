package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type GuideConf struct {
	Key string `yaml:"key"`
	Url string `yaml:"url"`
}

func (c *GuideConf) Build() *GuideConf {

	yamlFile, err := ioutil.ReadFile("yaml/guide.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
