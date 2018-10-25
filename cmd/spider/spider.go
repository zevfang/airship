package main

import (
	"github.com/zevfang/airship/cmd/spider/app"
	"github.com/zevfang/airship/pkg/adding"
	"github.com/zevfang/airship/pkg/storage/mongo"
	"log"
	"time"
)

func main() {
	w := app.NewCollector()
	w.QueryHot()

}

func Add() {
	storage, err := mongo.NewStorage()
	if err != nil {
		log.Fatal("init error.")
	}
	service := adding.NewService(storage)
	service.AddArticle(adding.Article{Title: "宫崎骏终身成就奖", Date: time.Now(), Content: "据悉，2018xxxxx"})
}
