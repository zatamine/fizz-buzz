package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServe(t *testing.T) {
	storage := NewInmemory()
	service := NewfizzBuzzService(storage)
	controller := NewController(service)
	router := NewGinRouter(controller)
	go func() {
		assert.NoError(t, router.Serve(":9000"))
	}()

	check := func() bool {
		if err := router.Serve(":9000"); err != nil {
			return true
		}
		return false
	}
	checkTimout := 500 * time.Millisecond
	checkPeriod := 2 * time.Millisecond
	assert.Eventually(t, check, checkTimout, checkPeriod)

	client := &http.Client{}

	resp1, err := client.Get("http://localhost:9000/fizz-buzz/stats")
	assert.NoError(t, err)
	defer resp1.Body.Close()
	assert.Equal(t, resp1.StatusCode, http.StatusOK)

	resp2, err := client.Get("http://localhost:9000/fizz-buzz/")
	assert.NoError(t, err)
	defer resp2.Body.Close()
	assert.Equal(t, resp2.StatusCode, http.StatusBadRequest)

	resp3, err := client.Get("http://localhost:9000/not-found/")
	assert.NoError(t, err)
	defer resp3.Body.Close()
	assert.Equal(t, resp3.StatusCode, http.StatusNotFound)
}
