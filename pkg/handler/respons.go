package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type error struct {
	Message string `json: "message"`
}

func NewErrorRespons(c *gin.Context, statusCode int, message string) {
	log.Fatal(message)
	c.AbortWithStatusJSON(statusCode, error{Message: message})

}
