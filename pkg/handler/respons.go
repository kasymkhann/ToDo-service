package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json: "message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Fatal(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})

}

type statusResponse struct {
	Status string `json:"status"`
}
