package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"telegraph-clone/internal/models"
	"telegraph-clone/internal/usecase"
)

type DataHandler struct {
	uc *usecase.DataUsecase
}

func NewDataHandler(uc *usecase.DataUsecase) *DataHandler {
	return &DataHandler{
		uc: uc,
	}
}

