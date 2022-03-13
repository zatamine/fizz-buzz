package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	tests := []struct {
		name        string
		inputParams FizzBuzzParams
		want        string
	}{
		{
			"With empty fizzbuzz params",
			FizzBuzzParams{},
			"6547cf7c-f553-512c-9c08-067c96057a9b",
		},
		{
			"With default fizzbuzz params",
			FizzBuzzParams{
				Int1:  3,
				Int2:  5,
				Limit: 100,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			"6ea78158-7b1d-5b30-9d27-912644a4e58b",
		},
		{
			"With min fizzbuzz params values",
			FizzBuzzParams{
				Int1:  1,
				Int2:  1,
				Limit: 1,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			"c9502844-788d-550f-9dd3-8daca38533d5",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fizzBuzz := tc.inputParams
			got := fizzBuzz.UUID()
			assert.Equal(t, tc.want, got)
		})
	}
}
