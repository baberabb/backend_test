package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func getHandler(c *gin.Context) {
	// Get query parameters from request
	long, _ := strconv.ParseFloat(c.Query("long"), 64)
	lat, _ := strconv.ParseFloat(c.Query("lat"), 64)
	distance, _ := strconv.ParseFloat(c.Query("distance"), 64)
	circle := false
	if c.Query("circle") == "true" {
		circle = true
	}

	// Connect to database
	sqlInfo := fmt.Sprintf("postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	db, err := sqlx.Connect("postgres", sqlInfo)
	if err != nil {
		fmt.Print(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Get data
	results, err := getdata(db, long, lat, distance, circle)
	if err != nil {
		print(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Encode as JSON
	c.JSON(http.StatusOK, results)
}
