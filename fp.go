package gfn

/* @example Map
gfn.Map([]int{1, 2, 3}, func(i int) string {
    return strconv.Itoa(i)
})
// []string{"1", "2", "3"}
*/

// Map returns a new array with the results of calling the mapper function on each element.
// @example
// gfn.Map([]int{1, 2, 3}, func(i int) string { return i+1 })
// // []int{2, 3, 4}
func Map[T any, R any](array []T, mapper func(T) R) []R {
	result := make([]R, len(array))
	for i, v := range array {
		result[i] = mapper(v)
	}
	return result
}

// Filter returns a new array containing elements of the original array
// that satisfy the provided function.
// @example
// array := []int{1, 2, 3, 4, 5, 6}
// gfn.Filter(array, func(i int) bool { return i%2 == 0 })
// // []int{2, 4, 6}
func Filter[T any](array []T, filter func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range array {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
