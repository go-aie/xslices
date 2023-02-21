package xslices

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func NumberToInt64[E Number](s []E) []int64 {
	var result []int64
	for _, x := range s {
		result = append(result, int64(x))
	}
	return result
}

func NumberToFloat64[E Number](s []E) []float64 {
	var result []float64
	for _, x := range s {
		result = append(result, float64(x))
	}
	return result
}

func Float64ToNumber[E Number](s []float64) []E {
	var result []E
	for _, x := range s {
		result = append(result, E(x))
	}
	return result
}
