package util

func IntersectFiles[T comparable](a, b []T, cmp func(T, T) bool) []T {
	res := make([]T, 0)
	for _, k := range a {
		if !compare(b, k, cmp) {
			res = append(res, k)
		}
	}
	return res
}

func compare[T any](a []T, b T, cmp func(T, T) bool) bool {
	for _, v := range a {
		if cmp(v, b) {
			return true
		}
	}
	return false
}
