package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGET(t *testing.T) {
	storage := NewInMemory()
	service := NewFizzBuzzService(storage)
	controller := NewController(service)
	router := NewGinRouter(controller)
	router.GET("/path", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/path", req.URL.Path)
	})
}
