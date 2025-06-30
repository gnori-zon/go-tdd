package slices

func Reduce[T any, R any](collection []T, fn func(R, T) R, initValue R) R {
	result := initValue
	for _, item := range collection {
		result = fn(result, item)
	}
	return result
}
