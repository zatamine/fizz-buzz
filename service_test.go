package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name   string
		params FizzBuzzParams
		want   *FizzBuzzOut
	}{
		{
			"With empty params",
			FizzBuzzParams{},
			&FizzBuzzOut{},
		},
		{
			"With default fizzbuzz params",
			FizzBuzzParams{
				Int1:  3,
				Int2:  5,
				Limit: 100,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			&FizzBuzzOut{
				Body: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz,31,32,fizz,34,buzz,fizz,37,38,fizz,buzz,41,fizz,43,44,fizzbuzz,46,47,fizz,49,buzz,fizz,52,53,fizz,buzz,56,fizz,58,59,fizzbuzz,61,62,fizz,64,buzz,fizz,67,68,fizz,buzz,71,fizz,73,74,fizzbuzz,76,77,fizz,79,buzz,fizz,82,83,fizz,buzz,86,fizz,88,89,fizzbuzz,91,92,fizz,94,buzz,fizz,97,98,fizz,buzz",
			},
		},
		{
			"With min fizzbuzz params values",
			FizzBuzzParams{
				Int1:  1,
				Int2:  1,
				Limit: 1,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			&FizzBuzzOut{
				Body: "fizzbuzz",
			},
		},
		{
			"With custom fizz buzz params",
			FizzBuzzParams{
				Int1:  3,
				Int2:  5,
				Limit: 100,
				Str1:  "foo",
				Str2:  "bar",
			},
			&FizzBuzzOut{
				Body: "1,2,foo,4,bar,foo,7,8,foo,bar,11,foo,13,14,foobar,16,17,foo,19,bar,foo,22,23,foo,bar,26,foo,28,29,foobar,31,32,foo,34,bar,foo,37,38,foo,bar,41,foo,43,44,foobar,46,47,foo,49,bar,foo,52,53,foo,bar,56,foo,58,59,foobar,61,62,foo,64,bar,foo,67,68,foo,bar,71,foo,73,74,foobar,76,77,foo,79,bar,foo,82,83,foo,bar,86,foo,88,89,foobar,91,92,foo,94,bar,foo,97,98,foo,bar",
			},
		},
		{
			"With customs params",
			FizzBuzzParams{
				Int1:  7,
				Int2:  21,
				Limit: 50,
				Str1:  "alice",
				Str2:  "bob",
			},
			&FizzBuzzOut{
				Body: "1,2,3,4,5,6,alice,8,9,10,11,12,13,alice,15,16,17,18,19,20,alicebob,22,23,24,25,26,27,alice,29,30,31,32,33,34,alice,36,37,38,39,40,41,alicebob,43,44,45,46,47,48,alice,50",
			},
		},
	}
	storage := NewInMemory()
	FBService := NewFizzBuzzService(storage)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FBService.FizzBuzz(tc.params)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestStatsService(t *testing.T) {
	testFizzBuzzParams := FizzBuzzParams{
		Int1:  3,
		Int2:  5,
		Limit: 100,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	testFooBarParams := FizzBuzzParams{
		Int1:  3,
		Int2:  5,
		Limit: 100,
		Str1:  "foo",
		Str2:  "bar",
	}
	testAliceBobParams := FizzBuzzParams{
		Int1:  3,
		Int2:  5,
		Limit: 100,
		Str1:  "alice",
		Str2:  "bob",
	}

	tests := []struct {
		name     string
		params1  FizzBuzzParams
		params2  FizzBuzzParams
		params3  FizzBuzzParams
		want     []FizzBuzzIn
		callNbr1 int
		callNbr2 int
		callNbr3 int
	}{
		{
			"With empty params",
			FizzBuzzParams{},
			FizzBuzzParams{},
			FizzBuzzParams{},
			[]FizzBuzzIn{},
			0,
			0,
			0,
		},
		{
			"With params and fizzBuzz params as most requested",
			testFizzBuzzParams,
			testFooBarParams,
			testAliceBobParams,
			[]FizzBuzzIn{
				{
					Hits:   5,
					Params: testFizzBuzzParams,
				},
				{
					Hits:   3,
					Params: testAliceBobParams,
				},
				{
					Hits:   2,
					Params: testFooBarParams,
				},
			},
			5,
			2,
			3,
		},
		{
			"With params and fooBar params as most requested",
			testFizzBuzzParams,
			testFooBarParams,
			testAliceBobParams,
			[]FizzBuzzIn{
				{
					Hits:   10,
					Params: testFooBarParams,
				},
				{
					Hits:   5,
					Params: testAliceBobParams,
				},
				{
					Hits:   3,
					Params: testFizzBuzzParams,
				},
			},
			3,
			10,
			5,
		},
	}

	for _, tc := range tests {
		storage := NewInMemory()
		FBService := NewFizzBuzzService(storage)
		t.Run(tc.name, func(t *testing.T) {
			for i := 1; i <= tc.callNbr1; i++ {
				FBService.FizzBuzz(tc.params1)
			}
			for i := 1; i <= tc.callNbr2; i++ {
				FBService.FizzBuzz(tc.params2)
			}
			for i := 1; i <= tc.callNbr3; i++ {
				FBService.FizzBuzz(tc.params3)
			}
			got := FBService.Stats()
			assert.Equal(t, tc.want, got)
		})
	}
}
