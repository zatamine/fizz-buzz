package main

import (
	"fmt"

	"github.com/google/uuid"
)

// ServiceError struc to format error message
type ServiceError struct {
	Message string `json:"message"`
	Level   string `json:"level"`
	Code    int    `json:"code"`
}

type FizzBuzzStorage struct {
	Data map[string]FizzBuzzIn
}

type FizzBuzzParams struct {
	Int1  int    `validate:"required,gte=1"`
	Int2  int    `validate:"required,gte=1"`
	Limit int    `validate:"required,gte=1"`
	Str1  string `validate:"required"`
	Str2  string `validate:"required"`
}

type FizzBuzzOut struct {
	Body string `json:"body"`
}

type FizzBuzzIn struct {
	Hits   int64
	Params FizzBuzzParams
}

func (f *FizzBuzzParams) UUID() string {
	const defaultUUID = "f47ac10b-58cc-0372-8567-0e02b2c3d479"

	inputParam := fmt.Sprintf("%v", f)
	newUUID := uuid.NewSHA1(uuid.MustParse(defaultUUID), []byte(inputParam))
	return newUUID.String()
}
