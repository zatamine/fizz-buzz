package main

import (
	"log"
)

func main() {
	storage := NewInmemory()
	service := NewfizzBuzzService(storage)
	controller := NewController(service)
	router := NewGinRouter(controller)

	if err := router.Serve(":8080"); err != nil {
		log.Fatalf("Cannot run HTTP server, %s", err)
	}
}
