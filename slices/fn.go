package slices

func Reduce[T any](collection []T, fn func(T, T) T, initValue T) T {
	result := initValue
	for _, item := range collection {
		result = fn(result, item)
	}
	return result
}
