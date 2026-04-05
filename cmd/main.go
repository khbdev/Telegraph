package main

import (
	"log"
	"telegraph-clone/internal/config"
	"telegraph-clone/internal/models"
	repository "telegraph-clone/internal/repostory"

	loadenv "telegraph-clone/pkg/loadEnv"
)


func main(){

	loadenv.Load()

	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	_ = db

	redis, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	_ = redis

	dataRepo := repository.NewDataRepository(db)
	urlRepo := repository.NewURLRepository(db)

	_ = dataRepo

   _ = urlRepo

   data := models.Data{
	Title: "Salom",
	YourName: "Azizbek Xasanov",
	YourStory: "Salom mening isimim Azizbek Xasanov",
   }

   dataRepo.Create(data)
}