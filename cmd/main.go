package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"telegraph-clone/internal/config"
	"telegraph-clone/internal/handler"
	"telegraph-clone/internal/usecase"

	repository "telegraph-clone/internal/repostory"

	loadenv "telegraph-clone/pkg/loadEnv"

	"github.com/gin-gonic/gin"
)


func main(){

go func() {
    http.ListenAndServe("localhost:6060", nil)
}()

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
    r := gin.Default()

  dataHandler := handler.NewDataHandler(dataUsecase)


  r.POST("/data", dataHandler.CreateData)


}