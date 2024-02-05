package typeutl

func Dereference[T any](t *T) T {
	if t == nil {
		return *new(T)
	}

	return *t
}
