package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/search", getHandler)
	err := router.Run(":8000")
	if err != nil {
		log.Print(err)
		return
	}
}
