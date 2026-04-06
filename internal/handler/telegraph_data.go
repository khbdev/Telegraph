package handler

import (
	"net/http"

	"telegraph-clone/internal/models"
	"telegraph-clone/internal/usecase"
	"telegraph-clone/pkg/response"

	"github.com/gin-gonic/gin"
)

type DataHandler struct {
	uc *usecase.DataUsecase
}

func NewDataHandler(uc *usecase.DataUsecase) *DataHandler {
	return &DataHandler{
		uc: uc,
	}
}



func (h *DataHandler) CreateData(c *gin.Context) {
	var input models.CreateDataInput

	if err := c.ShouldBindJSON(&input); err != nil {
		response.JSONError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	result, err := h.uc.CreateData(c.Request.Context(), input)
	if err != nil {
		response.JSONError(c, http.StatusInternalServerError, "something went wrong")
		return
	}


	response.JSONSuccess(c, http.StatusCreated, gin.H{
		"url":  result.URL,
		"data": result.Data,
	})
}

