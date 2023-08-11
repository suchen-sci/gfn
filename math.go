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

/* @example MaxBy
type Product struct {
	name   string
	amount int
}
products := []Product{
	{"apple", 10},
	{"banana", 20},
	{"orange", 30},
}
p := gfn.MaxBy(products, func(p Product) int {
	return p.amount
})  // {"orange", 30}
*/

// MaxBy returns the maximum value in the array, using the given function to transform values.
func MaxBy[T any, U Int | Uint | Float | ~string](array []T, fn func(T) U) T {
	if len(array) == 0 {
		panic("array is empty")
	}

	res := array[0]
	value := fn(res)
	for _, v := range array[1:] {
		current := fn(v)
		if current > value {
			res = v
			value = current
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

/* @example MinBy
type Product struct {
	name   string
	amount int
}
products := []Product{
	{"apple", 10},
	{"banana", 20},
	{"orange", 30},
}
p := gfn.MinBy(products, func(p Product) int {
	return p.amount
})  // {"apple", 10}
*/

// MinBy returns the maximum value in the array, using the given function to transform values.
func MinBy[T any, U Int | Uint | Float | ~string](array []T, fn func(T) U) T {
	if len(array) == 0 {
		panic("array is empty")
	}

	res := array[0]
	value := fn(res)
	for _, v := range array[1:] {
		current := fn(v)
		if current < value {
			res = v
			value = current
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

/* @example SumBy
type Product struct {
	name   string
	amount int
}
products := []Product{
	{"apple", 10},
	{"banana", 20},
	{"orange", 30},
}
gfn.SumBy(products, func(p Product) int {
	return p.amount
}) // 60
*/

// SumBy returns the sum of all values in the array after applying fn to each value.
func SumBy[T any, U Int | Uint | Float | ~string](array []T, fn func(T) U) U {
	if len(array) == 0 {
		panic("array is empty")
	}
	res := fn(array[0])
	for _, v := range array[1:] {
		res += fn(v)
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

// Mean returns the mean of all values in the array.
// @example
// gfn.Mean(1, 2, 3)               // 2.0
// gfn.Mean([]int{1, 2, 3, 4}...)  // 2.5
func Mean[T Int | Uint | Float](array ...T) float64 {
	if len(array) == 0 {
		panic("array is empty")
	}

	sum := 0.0
	for _, v := range array {
		sum += float64(v)
	}
	return sum / float64(len(array))
}

/* @example MeanBy
type Product struct {
	name string
	cost float64
}
products := []Product{
	{"apple", 1.5},
	{"banana", 2.5},
	{"orange", 3.5},
	{"lemon", 4.5},
}
gfn.MeanBy(products, func(p Product) float64 {
	return p.cost
})  // 3.0
*/

// MeanBy returns the mean of all values in the array after applying fn to each value.
func MeanBy[T any, U Int | Uint | Float](array []T, fn func(T) U) float64 {
	if len(array) == 0 {
		panic("array is empty")
	}

	sum := 0.0
	for _, v := range array {
		sum += float64(fn(v))
	}
	return sum / float64(len(array))
}
