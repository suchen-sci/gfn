// gfn is a Golang library that leverages generics to provide various methods.
package gfn

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
