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



func (h *DataHandler) CreateData(c *gin.Context) {
	var input models.CreateDataInput

	// 1. JSON parse
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	// 2. usecase chaqiramiz
	result, err := h.uc.CreateData(c.Request.Context(), input)
	if err != nil {
		// ❗ tizim xatoni chiqarmaymiz
		JSONError(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	// 3. response qaytaramiz
	JSONSuccess(c, http.StatusCreated, gin.H{
		"url":  result.URL,
		"data": result.Data,
	})
}