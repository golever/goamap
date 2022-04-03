package repo

import (
	"context"
	"goamap.mod/conf"
	"log"
)

type Response struct {
	Suggestion map[string]interface{}   `json:"suggestion"`
	Count      string                   `json:"Count"`
	Inode      string                   `json:"Infocode"`
	Poise      []map[string]interface{} `json:"Pois"`
	Status     string                   `json:"Status"`
	Info       string                   `json:"Info"`
}

func SavePoi(response *Response) {

	for i := range response.Poise {
		_, err := MGCollection(conf.Cfg.Mongo.Db, conf.Cfg.Mongo.Collection).
			InsertOne(context.TODO(), response.Poise[i])

		if err != nil {
			log.Printf("save data fail，data:%s", response.Poise[i][`name`])
		}
		log.Printf("sava data successful,data:%s", response.Poise[i][`name`])
	}
}
