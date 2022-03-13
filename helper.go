package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// show returns JSON HTTP response and status code
func show(response http.ResponseWriter, statusCode int, v interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	err, ok := v.(error)
	if ok {
		v = ServiceError{
			Message: err.Error(),
			Level:   "error",
			Code:    statusCode,
		}
	}
	json.NewEncoder(response).Encode(v)
}

// mapQueries map cqueries parameters with fizzBuzz struct
func mapQueries(request *http.Request) *FizzBuzzParams {
	queryValues := request.URL.Query()

	query1 := queryValues.Get("int1")
	int1, _ := parseInt(query1)

	query2 := queryValues.Get("int2")
	int2, _ := parseInt(query2)

	query3 := queryValues.Get("limit")
	limit, _ := parseInt(query3)

	fizzBuzz := FizzBuzzParams{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  queryValues.Get("str1"),
		Str2:  queryValues.Get("str2"),
	}
	return &fizzBuzz
}

// parseInt convetes string to interger
func parseInt(str string) (int, error) {
	return strconv.Atoi(str)
}
