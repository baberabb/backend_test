package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHandler(t *testing.T) {

	requestbody, err := json.Marshal(map[string]string{
		"long":     "0.0747",
		"lat":      "51.5250",
		"distance": "5000",
		"circle":   "true",
	})

	r := SetUpRouter()
	r.GET("/search", getHandler)
	req, _ := http.NewRequest("GET", "0.0.0.0/search", bytes.NewBuffer(requestbody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var testusers []Result
	err = json.Unmarshal(w.Body.Bytes(), &testusers)
	if err != nil {
		return
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, testusers)
}
