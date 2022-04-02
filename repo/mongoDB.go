package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goamap.mod/conf"
	"log"
	"time"
)

func MGCollection(db string, collect string) *mongo.Collection {

	client := options.Client().ApplyURI(conf.Cfg.Mongo.Url)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mc, err := mongo.Connect(ctx, client)
	if err != nil {
		log.Fatalf("connet mongo db error err:%s", err)
	}
	defer cancel()
	return mc.Database(db).Collection(collect)
}
