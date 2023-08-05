package gfn

import (
	"math"
)

// Max returns the maximum value in the array. For float64 arrays, please use MaxFloat64.
// NaN value in float64 arrays is not comparable to other values.
// Which means Max([math.NaN(), 0.5]) produces math.NaN(), but Max([0.5, math.NaN()]) produces 0.5.
// Since arrays with same elements but different order produce different results (inconsistent),
// this function does not support float64 arrays.
// @example
// gfn.Max([]int16{1, 5, 9, 10}...)  // 10
// gfn.Max("ab", "cd", "e")          // "e"
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
// @example
// gfn.MaxFloat64(1.1, math.NaN(), 2.2)                             // 2.2
// gfn.MaxFloat64([]float64{math.NaN(), math.NaN(), math.NaN()}...) // NaN
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
// @example
// gfn.Min(1.1, 2.2, 3.3)            // 1.1
// gfn.Min([]int16{1, 5, 9, 10}...)  // 1
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
// @example
// gfn.MinFloat64(1, -1, 10)                                   // -1
// gfn.MinFloat64([]float64{1.1, math.Inf(-1), math.NaN()}...) // math.Inf(-1)
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
// @example
// gfn.Sum([]int{1, 5, 9, 10}...)  // 25
// gfn.Sum(1.1, 2.2, 3.3)          // 6.6
// gfn.Sum("ab", "cd", "e")        // "abcde"
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
// @example
// gfn.Abs(-1)      // 1
// gfn.Abs(-100.99) // 100.99
func Abs[T Int | Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// DivMod returns quotient and remainder of a/b.
// @example
// gfn.DivMod(10, 3) // (3, 1)
func DivMod[T Int | Uint](a, b T) (T, T) {
	return a / b, a % b
}
