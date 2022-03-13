package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByParams(t *testing.T) {
	const host = "http://test"
	tests := []struct {
		name       string
		path       string
		statusCode int
	}{
		{
			"With empty path",
			"",
			400,
		},
		{
			"Without parameters",
			"/fizz-buzz",
			400,
		},
		{
			"With default parameters",
			"/fizz-buzz/?int1=3&int2=5&limit=100&str1=fizz&str2=buzz",
			200,
		},
		{
			"With bad parameters",
			"/fizz-buzz/?int2=5&limit=100&str1=fizz&str2=buzz",
			400,
		},
	}
	storage := NewInmemory()
	service := NewfizzBuzzService(storage)
	controller := NewController(service)
	for _, tc := range tests {
		url := fmt.Sprintf("%s%s", host, tc.path)
		request, err := http.NewRequest(http.MethodGet, url, nil)
		assert.NoError(t, err)
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(controller.GetByParams)
		handler.ServeHTTP(response, request)
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.statusCode, response.Code)
		})
	}
}

func TestStats(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       string
	}{
		{
			"Without stats parameters",
			200,
			"null\n",
		},
		{
			"With stats parameters",
			200,
			"[{\"Hits\":1,\"Params\":{\"Int1\":3,\"Int2\":5,\"Limit\":100,\"Str1\":\"fizz\",\"Str2\":\"buzz\"}}]\n",
		},
	}
	storage := NewInmemory()
	service := NewfizzBuzzService(storage)
	controller := NewController(service)
	for _, tc := range tests {
		if tc.want != "null\n" {
			callFizzBuzz(t, controller)
		}
		request, err := http.NewRequest(http.MethodGet, "", nil)
		assert.NoError(t, err)
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(controller.Stats)
		handler.ServeHTTP(response, request)
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.statusCode, response.Code)
			assert.Equal(t, tc.want, response.Body.String())
		})
	}
}

func callFizzBuzz(t *testing.T, controller IController) {
	t.Helper()
	request, err := http.NewRequest(http.MethodGet, "/fizz-buzz/?int1=3&int2=5&limit=100&str1=fizz&str2=buzz", nil)
	assert.NoError(t, err)
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetByParams)
	handler.ServeHTTP(response, request)
}
