package main

import (
	"fmt"
	"sort"
	"strings"
)

type FizzBuzzService struct {
	storage IStorage
}

type IFizzBuzzService interface {
	FizzBuzz(params FizzBuzzParams) *FizzBuzzOut
	Stats() []FizzBuzzIn
}

// NewfizzBuzzService a constructor to create a new instance of fizz-buzz service
func NewFizzBuzzService(storage IStorage) IFizzBuzzService {
	return &FizzBuzzService{
		storage: storage,
	}
}

func (s *FizzBuzzService) FizzBuzz(params FizzBuzzParams) *FizzBuzzOut {
	var result string
	const separator = ","
	for i := 1; i <= params.Limit; i++ {
		if i%params.Int1 == 0 && i%params.Int2 == 0 {
			result = fmt.Sprintf("%s%s%s%s", result, params.Str1, params.Str2, separator)
		} else if i%params.Int1 == 0 {
			result = fmt.Sprintf("%s%s%s", result, params.Str1, separator)
		} else if i%params.Int2 == 0 {
			result = fmt.Sprintf("%s%s%s", result, params.Str2, separator)
		} else {
			result = fmt.Sprintf("%s%d%s", result, i, separator)
		}
	}
	s.storage.Store(params)
	return &FizzBuzzOut{strings.TrimRight(result, separator)}
}

func (s *FizzBuzzService) Stats() []FizzBuzzIn {
	return s.sorted()
}

func (s *FizzBuzzService) sorted() []FizzBuzzIn {
	fizzBuzzList := s.storage.GetAll()
	sort.Slice(fizzBuzzList, func(i, j int) bool { return fizzBuzzList[i].Hits > fizzBuzzList[j].Hits })
	return fizzBuzzList
}
