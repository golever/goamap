package services

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"goamap.mod/db"
	"goamap.mod/request"
	"log"
	"net/http"
)

var (
	client     = db.GetMgoCli()
	err        error
	collection *mongo.Collection
	iResult    *mongo.InsertOneResult
)

type Response struct {
	Suggestion map[string]interface{}   `json:"suggestion"`
	Count      string                   `json:"Count"`
	Infocode   string                   `json:"Infocode"`
	Pois       []map[string]interface{} `json:"Pois"`
	Status     string                   `json:"Status"`
	Info       string                   `json:"Info"`
}

func KeyWords(c *gin.Context) {
	k := c.Param("keywords")
	t := c.Param("types")
	p := request.Params{Keywords: k, Types: t}

	content := request.Send(p)

	var data Response
	err := json.Unmarshal([]byte(content), &data)

	if err != nil {
		log.Printf("json format fail error info:%s\n", err.Error())
	}

	if data.Info == "OK" {
		collection = client.Database("goamap").Collection("poi")
		for i := range data.Pois {
			iResult, err = collection.InsertOne(context.TODO(), data.Pois[i])
			log.Printf("%s 保存成功\n", data.Pois[i][`name`])
		}
		c.JSON(http.StatusOK, "数据保存完成")
	}

}
