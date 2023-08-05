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
