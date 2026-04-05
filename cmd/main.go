package main

import (
	"log"
	"telegraph-clone/internal/config"
	"telegraph-clone/internal/usecase"

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

   ctx, cancel := con

   data := dataUsecase.CreateData()

}