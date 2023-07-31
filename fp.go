package gfn

// Map returns a new array with the results of calling the mapper function on each element.
func Map[T any, R any](array []T, mapper func(T) R) []R {
	result := make([]R, len(array))
	for i, v := range array {
		result[i] = mapper(v)
	}
	return result
}
