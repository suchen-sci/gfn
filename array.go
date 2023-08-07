// gfn is a Golang library that leverages generics to provide various methods.
package gfn

import "math/rand"

// Contains returns true if the array contains the value.
// @example
// gfn.Contains([]int{1, 2, 3}, 2)             // true
// gfn.Contains([]string{"a", "b", "c"}, "b")  // true
// gfn.Contains([]time.Duration{time.Second}, time.Second)  // true
func Contains[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// Range function returns a sequence of numbers, starting from start,
// and increments by 1, until end is reached (not included).
// @example
// gfn.Range(0, 7)    // []int{0, 1, 2, 3, 4, 5, 6}
// gfn.Range(3, 8)    // []int{3, 4, 3, 6, 7}
// gfn.Range(-10, -5) // []int{-10, -9, -8, -7, -6}
func Range[T Int | Uint](start, end T) []T {
	if start >= end {
		return []T{}
	}

	res := make([]T, end-start)
	for i := 0; i < len(res); i++ {
		res[i] = start + T(i)
	}
	return res
}

// RangeBy function returns a sequence of numbers, starting from start,
// and increments/decrements by step, until end is reached (not included).
// Zero step panics.
// @example
// gfn.RangeBy(0, 7, 1)   // []int{0, 1, 2, 3, 4, 5, 6}
// gfn.RangeBy(0, 8, 2)   // []int{0, 2, 4, 6}
// gfn.RangeBy(10, 0, -2) // []int{10, 8, 6, 4, 2}
func RangeBy[T Int | Uint](start, end, step T) []T {
	if step == 0 {
		panic("step must not be zero")
	}

	if start < end && step > 0 {
		res := make([]T, 0, (end-start)/step)
		for i := start; i < end; i += step {
			res = append(res, i)
		}
		return res
	}
	if start > end && step < 0 {
		res := make([]T, 0, (end-start)/step)
		for i := start; i > end; i += step {
			res = append(res, i)
		}
		return res
	}
	return []T{}
}

// Shuffle randomizes the order of elements by using Fisher–Yates algorithm
// @example
// array := []int{1, 2, 3, 4}
// gfn.Shuffle(array)
// // array: []int{2, 1, 4, 3} or other random order
func Shuffle[T any](array []T) {
	for i := range array {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

// Equal returns true if two arrays are equal.
// @example
// gfn.Equal([]int{1, 2, 3}, []int{1, 2, 3})                    // true
// gfn.Equal([]string{"a", "c", "b"}, []string{"a", "b", "c"})  // false
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aa := range a {
		if aa != b[i] {
			return false
		}
	}
	return true
}

// ToSet converts an array to a set.
// @example
// gfn.ToSet([]int{0, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5})
// // map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}}
func ToSet[T comparable](array []T) map[T]struct{} {
	res := make(map[T]struct{})
	for _, v := range array {
		res[v] = struct{}{}
	}
	return res
}

// IsSortedBy returns true if the array is sorted in the given order.
// The order function should return true if a1 is ok to be placed before a2.
// @example
// gfn.IsSortedBy([]int{2, 2, 1, 1, -1, -1}, func(a, b int) bool { return a >= b })
// // true
func IsSortedBy[T any](array []T, order func(a1, a2 T) bool) bool {
	for i := 0; i < len(array)-1; i++ {
		if !order(array[i], array[i+1]) {
			return false
		}
	}
	return true
}

// IsSorted returns true if the array is sorted in ascending order.
// @example
// gfn.IsSorted([]int{1, 2, 3, 4})  // true
func IsSorted[T Int | Uint | Float | ~string](array []T) bool {
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			return false
		}
	}
	return true
}

// Distribution returns a map of values and their counts.
// @example
// gfn.Distribution([]int{1, 2, 2, 2, 2})  // map[int]int{1: 1, 2: 4}
func Distribution[T comparable](array []T) map[T]int {
	res := make(map[T]int)
	for _, v := range array {
		res[v]++
	}
	return res
}

/* @example Zip
gfn.Zip([]int{1, 2, 3}, []string{"a", "b", "c"})
// []gfn.Pair[int, string]{
// 	{First: 1, Second: "a"},
// 	{First: 2, Second: "b"},
// 	{First: 3, Second: "c"}
// }
*/

// Zip returns a sequence of pairs built from the elements of two arrays.
func Zip[T, U any](a []T, b []U) []Pair[T, U] {
	l := Min(len(a), len(b))
	res := make([]Pair[T, U], l)
	for i := 0; i < l; i++ {
		res[i] = Pair[T, U]{a[i], b[i]}
	}
	return res
}

/* @example Unzip
pairs := []gfn.Pair[int, string]{
	{First: 1, Second: "a"},
	{First: 2, Second: "b"},
	{First: 3, Second: "c"},
}
gfn.Unzip(len(pairs), func(i int) (int, string) {
	return pairs[i].First, pairs[i].Second
})
// ([]int{1, 2, 3}, []string{"a", "b", "c"})
*/

// Unzip returns two arrays built from the elements of a sequence of pairs.
func Unzip[T, U any](n int, unzipFn func(i int) (T, U)) ([]T, []U) {
	if n < 0 {
		panic("negative length")
	}
	a := make([]T, n)
	b := make([]U, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = unzipFn(i)
	}
	return a, b
}

// Sample returns a random sample of n elements from an array.
// @example
// gfn.Sample([]int{1, 2, 3, 4, 5}, 3)  // []int{3, 1, 5} or other random choices.
func Sample[T any](array []T, n int) []T {
	if n < 0 {
		panic("negative length")
	}
	if n > len(array) {
		panic("sample size larger than array length")
	}
	indexes := Range(0, n)
	Shuffle(indexes)
	res := make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = array[indexes[i]]
	}
	return res
}

// Uniq returns an array with all duplicates removed.
// @example
// gfn.Uniq([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})  // []int{1, 2, 3, 4}
func Uniq[T comparable](array []T) []T {
	res := []T{}
	seen := make(map[T]struct{})
	for _, v := range array {
		if _, ok := seen[v]; !ok {
			res = append(res, v)
			seen[v] = struct{}{}
		}
	}
	return res
}

// Union returns an array with all duplicates removed from multiple arrays.
// @example
// gfn.Union([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
// // []int{1, 2, 3, 4, 5}
func Union[T comparable](arrays ...[]T) []T {
	res := []T{}
	seen := make(map[T]struct{})
	for _, array := range arrays {
		for _, v := range array {
			if _, ok := seen[v]; !ok {
				res = append(res, v)
				seen[v] = struct{}{}
			}
		}
	}
	return res
}

// Copy returns a new array that is a shallow copy of the original array.
// @example
// gfn.Copy([]int{1, 2, 3})  // []int{1, 2, 3}
//
// array := []int{1, 2, 3, 4, 5, 6}
// gfn.Copy(array[2:])
// // []int{3, 4, 5, 6}
func Copy[T any](array []T) []T {
	res := make([]T, len(array))
	copy(res, array)
	return res
}

// Diff returns a new array that is a copy of the original array,
// removing all occurrences of any item that also appear in others.
// The order is preserved from the original array.
// @example
// gfn.Diff([]int{1, 2, 3, 4}, []int{2, 4})  // []int{1, 3}
func Diff[T comparable](array []T, others ...[]T) []T {
	res := Copy(array)
	for _, other := range others {
		m := ToSet(other)
		res = Filter(res, func(v T) bool {
			_, ok := m[v]
			return !ok
		})
	}
	return res
}

// Fill sets all elements of an array to a given value.
// @example
// array := make([]bool, 5)
// Fill(array, true)
// // []bool{true, true, true, true, true}
//
// array2 := make([]int, 5)
// Fill(array2[2:], 100)
// // []int{0, 0, 100, 100, 100}
func Fill[T any](array []T, value T) {
	for i := range array {
		array[i] = value
	}
}

// Count returns the number of occurrences of a value in an array.
func Count[T comparable](array []T, value T) int {
	res := 0
	for _, v := range array {
		if v == value {
			res++
		}
	}
	return res
}

/* @example GroupBy
array := []int{1, 2, 3, 4, 5, 6, 7, 8}
groups := GroupBy(array, func(i int) string {
	if i%2 == 0 {
		return "even"
	}
	return "odd"
})
// map[string][]int{
// 	"even": []int{2, 4, 6, 8},
// 	"odd":  []int{1, 3, 5, 7},
// }
*/

// GroupBy generate a map of arrays by grouping the elements of an array
// according to a given function.
func GroupBy[T any, K comparable](array []T, groupFn func(T) K) map[K][]T {
	res := make(map[K][]T)
	for _, v := range array {
		k := groupFn(v)
		res[k] = append(res[k], v)
	}
	return res
}

// IndexOf returns the index of the first occurrence of a value in an array,
// or -1 if not found.
// @example
// gfn.IndexOf([]int{1, 2, 3, 4}, 3)  // 2
// gfn.IndexOf([]int{1, 2, 3, 4}, 5)  // -1
func IndexOf[T comparable](array []T, value T) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of a value in an array,
// or -1 if not found.
// @example
// gfn.LastIndexOf([]int{3, 3, 3, 4}, 3)  // 2
// gfn.LastIndexOf([]int{1, 2, 3, 4}, 5)  // -1
func LastIndexOf[T comparable](array []T, value T) int {
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == value {
			return i
		}
	}
	return -1
}

// Reverse reverses an array in place.
// @example
// array := []int{1, 2, 3, 4}
// gfn.Reverse(array)
// // []int{4, 3, 2, 1}
func Reverse[T any](array []T) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

/* @example All
gfn.All([]int{1, 2, 3, 4}, func(i int) bool {
	return i > 0
}
// true
*/

// All returns true if all elements in an array pass a given test.
func All[T any](array []T, fn func(T) bool) bool {
	for _, v := range array {
		if !fn(v) {
			return false
		}
	}
	return true
}

/* @example Any
gfn.Any([]int{1, 2, 3, 4}, func(i int) bool {
	return i > 3
}
// true
*/

// Any returns true if at least one element in an array passes a given test.
func Any[T any](array []T, fn func(T) bool) bool {
	for _, v := range array {
		if fn(v) {
			return true
		}
	}
	return false
}
