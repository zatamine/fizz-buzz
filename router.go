package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type router struct {
	ginDispatcher *gin.Engine
	controller    IController
}

type IRouter interface {
	GET(path string, f func(w http.ResponseWriter, req *http.Request))
	Serve(addr string) error
}

// NewGinRouter a constructor to create new instance of Gin-gonic HTTP server
func NewGinRouter(controller IController) IRouter {
	gin.SetMode(gin.ReleaseMode)
	ginDispatcher := gin.New()
	ginDispatcher.Use(
		gin.Recovery(), // Recover from any panic and return HTTP 500 status code
	)
	router := router{
		ginDispatcher: ginDispatcher,
		controller:    controller,
	}
	paths(&router)
	return &router
}

func (r *router) GET(path string, f func(w http.ResponseWriter, req *http.Request)) {
	r.ginDispatcher.GET(path, gin.WrapF(f))
}

func (r *router) Serve(addr string) error {
	log.Printf("GIN HTTP server running on port '%v'\n", addr)
	return r.ginDispatcher.Run(addr)
}

func paths(r *router) {
	r.GET("/fizz-buzz/stats", r.controller.Stats)
	r.GET("/fizz-buzz", r.controller.GetByParams)
}
