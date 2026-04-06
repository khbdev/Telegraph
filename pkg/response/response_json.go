package response

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}


func JSONSuccess(c *gin.Context, status int, data interface{}) {
	c.JSON(status, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func JSONError(c *gin.Context, status int, message string) {
	c.JSON(status, ErrorResponse{
		Success: false,
		Error:   message,
	})
}