package main

import (
	"L0/internal/cache"
	"L0/internal/database"
	"L0/internal/nats"
	"L0/internal/server"
	"fmt"
)

func init() {
	var db database.Database
	cashe := make(cache.TypeCache)

	db.Connect()
	orders := db.GetAllOrders()
	cashe = cashe.Warming(orders)
	cacheSubscribe(&cashe)
	fmt.Println("init")
	DbSubscribe(&db)
	fmt.Println("init2")

}

func main() {
	server.RunServer()
}

func DbSubscribe(db *database.Database) {
	var srvNats nats.Service
	err := srvNats.Connect("consumer_db")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = srvNats.Subscribe("upload_consumer", db)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func cacheSubscribe(c *cache.TypeCache) {
	var srvNats nats.Service
	err := srvNats.Connect("consumer_cache")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = srvNats.Subscribe("upload_consumer", c)
	if err != nil {
		fmt.Println(err)
		return
	}
}
// nats-server -c /путь/к/nats-server.conf