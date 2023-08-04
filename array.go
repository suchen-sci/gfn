// gfn is a Golang library that leverages generics to provide various methods.
package gfn

import "math/rand"

// Contains returns true if the array contains the value.
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
func Shuffle[T any](array []T) {
	for i := range array {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

// Equal returns true if two arrays are equal.
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
func ToSet[T comparable](array []T) map[T]struct{} {
	res := make(map[T]struct{})
	for _, v := range array {
		res[v] = struct{}{}
	}
	return res
}

// IsSortedBy returns true if the array is sorted in the given order.
// The order function should return true if a1 is ok to be placed before a2.
func IsSortedBy[T any](array []T, order func(a1, a2 T) bool) bool {
	for i := 0; i < len(array)-1; i++ {
		if !order(array[i], array[i+1]) {
			return false
		}
	}
	return true
}

// Distribution returns a map of values and their counts.
func Distribution[T comparable](array []T) map[T]int {
	res := make(map[T]int)
	for _, v := range array {
		res[v]++
	}
	return res
}
