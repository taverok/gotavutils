package sliceutl

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
