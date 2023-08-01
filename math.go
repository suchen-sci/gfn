package gfn

import "math"

// Max returns the maximum value in the array.
// Be careful when using this function for float64 arrays with NaN values.
// NaN values is not comparable to other float64. NaN > x is false. NaN < x is false.
// which means Max([math.NaN(), 0.5]) is math.NaN(), but Max([0.5, math.NaN()]) is 0.5.
func Max[T Int | Uint | Float | ~string](array ...T) T {
	res := array[0]
	for _, v := range array {
		if v > res {
			res = v
		}
	}
	return res
}

// MaxNotNaN returns the maximum not NaN value in the array.
func MaxNotNaN(array ...float64) float64 {
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

// Min returns the minimum value in the array.
// Be careful when using this function for float64 arrays with NaN values.
// NaN values is not comparable to other float64. NaN > x is false. NaN < x is false.
// which means Min([math.NaN(), 0.5]) is math.NaN(), but Min([0.5, math.NaN()]) is 0.5.
func Min[T Int | Uint | Float | ~string](array ...T) T {
	res := array[0]
	for _, v := range array {
		if v < res {
			res = v
		}
	}
	return res
}

// MinNotNaN returns the minimum not NaN value in the array.
func MinNotNaN(array ...float64) float64 {
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
// NaN values is not comparable to other float64. NaN + x is NaN.
// which means Sum([math.NaN(), 0.5]) is math.NaN(). Sum(math.Inf(1), math.Inf(-1)) is math.NaN() too.
func Sum[T Int | Uint | Float | ~string | Complex](array ...T) T {
	res := array[0]
	for i, v := range array {
		if i > 0 {
			res += v
		}
	}
	return res
}

// SumNotNaN returns the sum not NaN values in the array.
func SumNotNaN(array ...float64) float64 {
	if len(array) == 0 {
		panic("array is empty")
	}

	// find the first not NaN value
	i := 0
	for ; i < len(array); i++ {
		if !math.IsNaN(array[i]) {
			break
		}
	}

	// all values are NaN
	if i == len(array) {
		return math.NaN()
	}

	res := array[i]
	for _, v := range array[i+1:] {
		if math.IsNaN(v) {
			continue
		}
		res += v
	}
	return res
}
