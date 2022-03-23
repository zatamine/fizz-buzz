package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapQueries(t *testing.T) {
	const host = "http://localhost"
	tests := []struct {
		name      string
		inputPath string
		want      *FizzBuzzParams
	}{
		{
			"With empty input",
			"",
			&FizzBuzzParams{},
		},
		{
			"With good path input",
			"/?int1=3&int2=5&limit=100&str1=fizz&str2=buzz",
			&FizzBuzzParams{3, 5, 100, "fizz", "buzz"},
		},
		{
			"With bad int1 query parameter",
			"/?int1=@&int2=5&limit=100&str1=fizz&str2=buzz",
			&FizzBuzzParams{0, 5, 100, "fizz", "buzz"},
		},
		{
			"Without int1 query parameter",
			"/?int2=5&limit=100&str1=fizz&str2=buzz",
			&FizzBuzzParams{0, 5, 100, "fizz", "buzz"},
		},
		{
			"Without int2 query parameter",
			"/?int1=3&limit=100&str1=fizz&str2=buzz",
			&FizzBuzzParams{3, 0, 100, "fizz", "buzz"},
		},
		{
			"Without limit query parameter",
			"/?int1=3&int2=5&str1=fizz&str2=buzz",
			&FizzBuzzParams{3, 5, 0, "fizz", "buzz"},
		},
		{
			"Without str1 query parameter",
			"/?int1=3&int2=5&limit=100&str2=buzz",
			&FizzBuzzParams{3, 5, 100, "", "buzz"},
		},
		{
			"Without str2 query parameter",
			"/?int1=3&int2=5&limit=100&str1=fizz",
			&FizzBuzzParams{3, 5, 100, "fizz", ""},
		},
		{
			"Without str1 and str2 queries parameters",
			"/?int1=3&int2=5&limit=100",
			&FizzBuzzParams{3, 5, 100, "", ""},
		},
		{
			"Without int1 and int2 queries parameters",
			"/?limit=100&str1=fizz&str2=buzz",
			&FizzBuzzParams{0, 0, 100, "fizz", "buzz"},
		},
		{
			"Without int1, int2 and limit  queries parameters",
			"/?str1=fizz&str2=buzz",
			&FizzBuzzParams{0, 0, 0, "fizz", "buzz"},
		},
		{
			"Without int1, int2, limit and str1 queries parameters",
			"/?str2=buzz",
			&FizzBuzzParams{0, 0, 0, "", "buzz"},
		},
	}
	for _, tc := range tests {
		url := fmt.Sprintf("%s%s", host, tc.inputPath)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		assert.NoError(t, err)
		t.Run(tc.name, func(t *testing.T) {
			got := mapQueries(req)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		want     int
		expError bool
	}{
		{
			"With empty input",
			"",
			0,
			true,
		},
		{
			"With good input string",
			"10",
			10,
			false,
		},
		{
			"With bad input string",
			"@",
			0,
			true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseInt(tc.input)
			if tc.expError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
