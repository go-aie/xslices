package xslices

func Repeat[E any](s []E, n int) []E {
	if n <= 0 {
		return nil
	}
	result := make([]E, 0, n*len(s))
	for i := 0; i < n; i++ {
		result = append(result, s...)
	}
	return result
}

func Concat[E any](ss ...[]E) []E {
	var result []E
	for _, s := range ss {
		result = append(result, s...)
	}
	return result
}

func Indices[E comparable](s []E, v E) []int {
	var indices []int
	for i, x := range s {
		if x == v {
			indices = append(indices, i)
		}
	}
	return indices
}
