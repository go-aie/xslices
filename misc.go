package xslices

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Range(start, end int) []int {
	var result []int
	for i := start; i < end; i++ {
		result = append(result, i)
	}
	return result
}
