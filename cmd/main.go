package main

import (
	"log"
	"telegraph-clone/internal/config"
	loadenv "telegraph-clone/pkg/loadEnv"
)


func main(){

	loadenv.Load()

	db := config.ConnectDB()
	_ = db

	redis, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	_ = redis
}