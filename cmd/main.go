package main

import (
	// "context"
	// "encoding/json"
	// "fmt"
	"context"
	"encoding/json"
	// "encoding/json"
	"fmt"
	"log"
	"telegraph-clone/internal/config"
	"time"

	// "telegraph-clone/internal/models"
	// "telegraph-clone/internal/usecase"
	// "time"

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

 


}