package gfn

import (
	"math"
)

// Max returns the maximum value in the array. For float64 arrays, please use MaxFloat64.
// NaN value in float64 arrays is not comparable to other values.
// Which means Max([math.NaN(), 0.5]) produces math.NaN(), but Max([0.5, math.NaN()]) produces 0.5.
// Since arrays with same elements but different order produce different results (inconsistent),
// this function does not support float64 arrays.
func Max[T Int | Uint | ~float32 | ~string](array ...T) T {
	if len(array) == 0 {
		panic("array is empty")
	}

	res := array[0]
	for _, v := range array {
		if v > res {
			res = v
		}
	}
	return res
}

// MaxFloat64 returns the maximum value in the array. NaN values are skipped.
func MaxFloat64(array ...float64) float64 {
	if len(array) == 0 {
		panic("array is empty")
	}

	res := array[0]
	for _, v := range array {
		if math.IsNaN(v) {
			continue
		}
		if math.IsNaN(res) || v > res {
			res = v
		}
	}
	return res
}

// Min returns the minimum value in the array. For float64 arrays, please use MinFloat64.
// More details in Max.
func Min[T Int | Uint | ~float32 | ~string](array ...T) T {
	if len(array) == 0 {
		panic("array is empty")
	}

	res := array[0]
	for _, v := range array {
		if v < res {
			res = v
		}
	}
	return res
}

// MinFloat64 returns the minimum value in the array. NaN values are skipped.
func MinFloat64(array ...float64) float64 {
	if len(array) == 0 {
		panic("array is empty")
	}

	res := array[0]
	for _, v := range array {
		if math.IsNaN(v) {
			continue
		}
		if math.IsNaN(res) || v < res {
			res = v
		}
	}
	return res
}

// Sum returns the sum of all values in the array.
// Be careful when using this function for float64 arrays with NaN and Inf values.
// Sum([math.NaN(), 0.5]) produces math.NaN(). Sum(math.Inf(1), math.Inf(-1)) produces math.NaN() too.
func Sum[T Int | Uint | Float | ~string | Complex](array ...T) T {
	if len(array) == 0 {
		panic("array is empty")
	}

	res := array[0]
	for i, v := range array {
		if i > 0 {
			res += v
		}
	}
	return res
}

// Abs returns the absolute value of x.
func Abs[T Int | Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// DivMod returns quotient and remainder of a/b.
func DivMod[T Int | Uint](a, b T) (T, T) {
	return a / b, a % b
}

// IsSorted returns true if the array is sorted in ascending order.
func IsSorted[T Int | Uint | Float | ~string](array []T) bool {
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			return false
		}
	}
	return true
}
