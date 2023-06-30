package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/search", getHandler)
	log.Println("Server running on port 8000")
	err := router.Run(":8000")
	if err != nil {
		fmt.Print(err)
		return
	}
}
