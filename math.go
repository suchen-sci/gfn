package gfn

/* @example Max
gfn.Max([]int16{1, 5, 9, 10}...)  // 10
gfn.Max("ab", "cd", "e")          // "e"

gfn.Max(1.1, math.NaN(), 2.2)                             // 2.2
gfn.Max([]float64{math.NaN(), math.NaN(), math.NaN()}...) // NaN
*/

// Max returns the maximum value in the array. For float64 arrays, NaN values are skipped.
func Max[T Int | Uint | Float | ~string](array ...T) T {
	if len(array) == 0 {
		panic("array is empty")
	}

	res := array[0]
	for _, v := range array {
		if isNaN(v) {
			continue
		}
		if isNaN(res) || v > res {
			res = v
		}
	}
	return res
}

// isNaN reports whether input is an IEEE 754 "not-a-number" value.
func isNaN[T Int | Uint | Float | ~string](x T) bool {
	// IEEE 754 says that only NaNs satisfy x != x.
	return x != x
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

/* @example Min
gfn.Min(1.1, 2.2, 3.3)            // 1.1
gfn.Min([]int16{1, 5, 9, 10}...)  // 1

gfn.Min(1, -1, 10)                                   // -1
gfn.Min([]float64{1.1, math.Inf(-1), math.NaN()}...) // math.Inf(-1)
*/

// Min returns the minimum value in the array. For float64 arrays, NaN values are skipped.
func Min[T Int | Uint | Float | ~string](array ...T) T {
	if len(array) == 0 {
		panic("array is empty")
	}

	res := array[0]
	for _, v := range array {
		if isNaN(v) {
			continue
		}
		if isNaN(res) || v < res {
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

// MinBy returns the minimum value in the array, using the given function to transform values.
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

/* @example Sum
gfn.Sum([]int{1, 5, 9, 10}...)  // 25
gfn.Sum(1.1, 2.2, 3.3)          // 6.6
gfn.Sum("ab", "cd", "e")        // "abcde"
*/

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

/* @example Abs
gfn.Abs(-1)      // 1
gfn.Abs(-100.99) // 100.99
*/

// Abs returns the absolute value of x.
func Abs[T Int | Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

/* @example DivMod
gfn.DivMod(10, 3) // (3, 1)
*/

// DivMod returns quotient and remainder of a/b.
func DivMod[T Int | Uint](a, b T) (T, T) {
	return a / b, a % b
}

/* @example Mean
gfn.Mean(1, 2, 3)               // 2.0
gfn.Mean([]int{1, 2, 3, 4}...)  // 2.5
*/

// Mean returns the mean of all values in the array.
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

/* @example MinMax
gfn.MinMax(1, 5, 9, 10)  // 1, 10

gfn.MinMax(math.NaN(), 1.85, 2.2) // 1.85, 2.2
gfn.MinMax(math.NaN(), math.NaN(), math.NaN()) // NaN, NaN
*/

// MinMax returns the minimum and maximum value in the array. For float64 arrays, please use MinMaxFloat64.
func MinMax[T Int | Uint | Float | ~string](array ...T) (T, T) {
	if len(array) == 0 {
		panic("array is empty")
	}

	minimum := array[0]
	maximum := array[0]
	for _, v := range array {
		if isNaN(v) {
			continue
		}
		if isNaN(minimum) || v < minimum {
			minimum = v
		}
		if isNaN(maximum) || v > maximum {
			maximum = v
		}
	}
	return minimum, maximum
}

/* @example MinMaxBy
type Product struct {
	name   string
	amount int
}
products := []Product{
	{"banana", 20},
	{"orange", 30},
	{"apple", 10},
	{"grape", 50},
	{"lemon", 40},
}
gfn.MinMaxBy(products, func(p Product) int {
	return p.amount
}) // {"apple", 10}, {"grape", 50}
*/

// MinMaxBy returns the minimum and maximum value in the array, using the given function to transform values.
func MinMaxBy[T any, U Int | Uint | Float | ~string](array []T, fn func(T) U) (T, T) {
	if len(array) == 0 {
		panic("array is empty")
	}

	minimum := array[0]
	minValue := fn(minimum)
	maximum := minimum
	maxValue := minValue
	for _, v := range array[1:] {
		current := fn(v)
		if current < minValue {
			minimum = v
			minValue = current
		}
		if current > maxValue {
			maximum = v
			maxValue = current
		}
	}
	return minimum, maximum
}

/* @example Mode
gfn.Mode([]int{1, 1, 5, 5, 5, 2, 2})) // 5
*/

// Mode returns the most frequent value in the array.
func Mode[T comparable](array []T) T {
	if len(array) == 0 {
		panic("array is empty")
	}

	value := array[0]
	count := 1
	seen := make(map[T]int)

	for _, v := range array {
		seen[v]++
		if seen[v] > count {
			value = v
			count = seen[v]
		}
	}
	return value
}

/* @example ModeBy
type Product struct {
	name   string
	amount int
}
products := []Product{
	{"banana", 20},
	{"banana", 20},
	{"apple", 10},
}
gfn.ModeBy(products, func(p Product) int {
	return p.amount
}) // {"banana", 20}
*/

// ModeBy returns the most frequent value in the array, using the given function to transform values.
func ModeBy[T any, U comparable](array []T, fn func(T) U) T {
	if len(array) == 0 {
		panic("array is empty")
	}

	value := array[0]
	count := 1
	seen := make(map[U]int)

	for _, v := range array {
		current := fn(v)
		seen[current]++
		if seen[current] > count {
			value = v
			count = seen[current]
		}
	}
	return value
}
