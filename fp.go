package gfn

// Map returns a new array with the results of calling the mapper function on each element.
func Map[T any, R any](array []T, mapper func(T) R) []R {
	result := make([]R, len(array))
	for i, v := range array {
		result[i] = mapper(v)
	}
	return result
}

// Filter returns a new array containing elements of the original array
// that satisfy the provided function.
func Filter[T any](array []T, filter func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range array {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
