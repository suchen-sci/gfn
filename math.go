package gfn

import (
	"errors"
	"math"
)

var ErrEmptyArray = errors.New("array is empty")

// Max returns the maximum value in the array. For float64 arrays, please use MaxFloat64.
// NaN value in float64 arrays is not comparable to other values.
// Which means Max([math.NaN(), 0.5]) produces math.NaN(), but Max([0.5, math.NaN()]) produces 0.5.
// Since arrays with same elements but different order produce different results (inconsistent),
// this function does not support float64 arrays.
func Max[T Int | Uint | ~float32 | ~string](array ...T) T {
	if len(array) == 0 {
		panic(ErrEmptyArray)
	}

	res := array[0]
	for _, v := range array {
		if v > res {
			res = v
		}
	}
	return res
}

// MaxFloat64 returns the maximum value in the array. NaN value is skipped if skipNaN is true.
// If skipNaN is false, function returns NaN if NaN is present in the array.
// Even if skipNaN is true, function returns NaN if all values are NaN.
func MaxFloat64(skipNaN bool, array ...float64) float64 {
	if len(array) == 0 {
		panic(ErrEmptyArray)
	}

	res := array[0]
	for _, v := range array {
		if math.IsNaN(v) {
			if skipNaN {
				continue
			}
			return math.NaN()
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
		panic(ErrEmptyArray)
	}

	res := array[0]
	for _, v := range array {
		if v < res {
			res = v
		}
	}
	return res
}

// MinFloat64 returns the minimum value in the array. NaN value is skipped if skipNaN is true.
// More details in MaxFloat64.
func MinFloat64(skipNaN bool, array ...float64) float64 {
	if len(array) == 0 {
		panic(ErrEmptyArray)
	}

	res := array[0]
	for _, v := range array {
		if math.IsNaN(v) {
			if skipNaN {
				continue
			}
			return math.NaN()
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
		panic(ErrEmptyArray)
	}

	res := array[0]
	for i, v := range array {
		if i > 0 {
			res += v
		}
	}
	return res
}

// SumFloat64 returns the sum of all values in the array. NaN value is skipped if skipNaN is true.
// If skipNaN is false, function returns NaN if NaN is present in the array.
// Even if skipNaN is true, function returns NaN if all values are NaN or SumFloat64(math.Inf(1), math.Inf(-1)).
func SumFloat64(skipNaN bool, array ...float64) float64 {
	if len(array) == 0 {
		panic(ErrEmptyArray)
	}

	i := 0
	for ; i < len(array); i++ {
		if math.IsNaN(array[i]) {
			if !skipNaN {
				return math.NaN()
			}
		} else {
			break
		}
	}
	if i == len(array) {
		return math.NaN()
	}
	res := array[i]
	for _, v := range array[i+1:] {
		if math.IsNaN(v) {
			if skipNaN {
				continue
			}
			return math.NaN()
		}
		res += v
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
