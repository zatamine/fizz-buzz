package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// I'm ok with not testing this call
	os.Exit(realMain(os.Stdout))
}
func realMain(out io.Writer) int {

	storage := NewInmemory()
	service := NewfizzBuzzService(storage)
	controller := NewController(service)
	router := NewGinRouter(controller)

	if err := router.Serve(":8080"); err != nil {
		log.Printf("Cannot run HTTP server, %s", err)
		return 1
	}
	return 0
}
