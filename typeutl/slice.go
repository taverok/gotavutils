package typeutl

func Transform[T, R any](ee []T, f func(T) R) []R {
	var result []R
	for _, e := range ee {
		result = append(result, f(e))
	}
	return result
}

func Filter[T any](ee []T, f func(T) bool) []T {
	var result []T
	for _, e := range ee {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

func ToMap[K comparable, V any](ee []V, keyFunc func(V) K) map[K]V {
	result := make(map[K]V)
	for _, e := range ee {
		result[keyFunc(e)] = e
	}
	return result
}

func Intersect[T comparable](a, b []T) []T {
	aMap := make(map[T]bool)
	var result []T

	for _, e := range a {
		aMap[e] = true
	}

	for _, e := range b {
		if aMap[e] {
			result = append(result, e)
		}
	}

	return result
}
