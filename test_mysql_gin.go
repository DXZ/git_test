package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()
	router.GET("/test", func(c *gin.Context) {
		fmt.Println("test")
		c.String(http.StatusOK, "It works!")
	})
	router.Run(":9998")
}
