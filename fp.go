package gfn

/* @example Map
gfn.Map([]int{1, 2, 3}, func(i int) string {
	return strconv.Itoa(i)
})
// []string{"1", "2", "3"}
*/

// Map returns a new array with the results of calling the mapper function on each element.
// No MapKV because I don't know what to return, an array or a map? Instead, use ForEachKV.
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

/* @example FilterKV
m := map[int]string{1: "a", 2: "b", 3: "c"}
gfn.FilterKV(m, func(k int, v string) bool {
	return k == 1 || v == "c"
})
// map[int]string{1: "a", 3: "c"}
*/

// FilterKV returns a new map containing elements of the original map
// that satisfy the provided function.
func FilterKV[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	return Select(m, fn)
}

/* @example Reduce
gfn.Reduce([]int{1, 2, 3}, 0, func(a, b int) int {
	return a + b
})
// 6
*/

// Reduce executes a reducer function on each element of the array,
// resulting in a single output value.
func Reduce[T any, R any](array []T, initialValue R, reducer func(R, T) R) R {
	result := initialValue
	for _, v := range array {
		result = reducer(result, v)
	}
	return result
}

/* @example ReduceKV
m := map[string]int{"a": 1, "b": 2, "c": 3}
total := gfn.ReduceKV(m, 0, func(value int, k string, v int) int {
	return value + v
})
// 6
*/

// ReduceKV executes a reducer function on each element of the map,
// resulting in a single output value.
func ReduceKV[K comparable, V any, R any](m map[K]V, initialValue R, reducer func(R, K, V) R) R {
	result := initialValue
	for k, v := range m {
		result = reducer(result, k, v)
	}
	return result
}
