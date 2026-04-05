package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"telegraph-clone/internal/config"
	"telegraph-clone/internal/models"
	"telegraph-clone/internal/usecase"
	"time"

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

   dataUsecase := usecase.NewDataUsecase(dataRepo, urlRepo)

   _ = dataUsecase

   ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
   defer cancel()

   req := models.CreateDataInput{
	Title: "Salom",
	YourName: "Azizbek Xasanov",
	YourStory: "Salom Mening isimim Azizbek",
   }
   data, err := dataUsecase.CreateData(ctx, req)
if err != nil {
	log.Fatal(err)
}

jsonData, err := json.MarshalIndent(data, "", "  ")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(jsonData))



}