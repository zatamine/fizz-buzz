package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type controller struct {
	service  IFizzBuzzService
	validate *validator.Validate
}
type IController interface {
	GetByParams(response http.ResponseWriter, request *http.Request)
	Stats(response http.ResponseWriter, request *http.Request)
}

func NewController(service IFizzBuzzService) IController {
	validate := validator.New()
	return &controller{
		service:  service,
		validate: validate,
	}
}

func (c *controller) GetByParams(response http.ResponseWriter, request *http.Request) {
	fizzBuzz := mapQueries(request)
	if err := c.validate.Struct(fizzBuzz); err != nil {
		log.Println(err)
		show(response, http.StatusBadRequest, err)
		return
	}
	fizzBuzzsOut := c.service.FizzBuzz(*fizzBuzz)

	show(response, http.StatusOK, fizzBuzzsOut)
}

func (c *controller) Stats(response http.ResponseWriter, request *http.Request) {
	fizzBuzzStats := c.service.Stats()
	show(response, http.StatusOK, fizzBuzzStats)
}
