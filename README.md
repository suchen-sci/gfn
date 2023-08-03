# gfn <!-- omit in toc -->
`gfn` is a Golang library that leverages generics to provide various methods, including common functional programming techniques such as `Map`, `Reduce`, and `Filter`, along with other utilities like `Contains`, `Keys`, etc.

- [Installation](#installation)
- [Usage](#usage)
- [Type](#type)
- [Array](#array)
  - [gfn.Contains](#gfncontains)
  - [gfn.Range](#gfnrange)
  - [gfn.RangeBy](#gfnrangeby)
  - [gfn.Map](#gfnmap)
  - [gfn.Filter](#gfnfilter)
- [Map](#map)
- [Math](#math)
  - [gfn.Abs](#gfnabs)
  - [gfn.Max](#gfnmax)
  - [gfn.MaxFloat64](#gfnmaxfloat64)
  - [gfn.Min](#gfnmin)
  - [gfn.MinFloat64](#gfnminfloat64)
  - [gfn.Sum](#gfnsum)
  - [gfn.SumFloat64](#gfnsumfloat64)


## Installation
```
go get github.com/suchen-sci/gfn
```

## Usage 
```
import "github.com/suchen-sci/gfn"
```

## Type
```go
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}
```

## Array

### gfn.Contains

```go
func Contains[T comparable](array []T, value T) bool
```

Contains returns true if the array contains the value.

```go
gfn.Contains([]int{1, 2, 3}, 2)               // true
gfn.Contains([]string{"a", "b", "c"}, "b")    // true
```

### gfn.Range

```go
func Range[T Int | Uint](start, end, step T) []T 
```

Range function returns a sequence of numbers, starting from start, and increments by 1, until end is reached (not included).

```go
gfn.Range(0, 7)   // []int{0, 1, 2, 3, 4, 5, 6}
gfn.Range(3, 8)   // []int{3, 4, 3, 6, 7}
gfn.Range(-10, -5) // []int{-10, -9, -8, -7, -6}
```

### gfn.RangeBy

```go
func RangeBy[T Int | Uint](start, end, step T) []T 
```

RangeBy function returns a sequence of numbers, starting from start, and increments/decrements by step, until end is reached. Zero step panics (not included).

```go
gfn.RangeBy(0, 7, 1)   // []int{0, 1, 2, 3, 4, 5, 6}
gfn.RangeBy(0, 8, 2)   // []int{0, 2, 4, 6}
gfn.RangeBy(10, 0, -2) // []int{10, 8, 6, 4, 2}
```

### gfn.Map

```go
func Map[T any, R any](array []T, mapper func(T) R) []R
```

Map returns a new array with the results of calling the mapper function on each element.

```go
// map int to string
gfn.Map([]int{1, 2, 3}, func(i int) string { 
    return strconv.Itoa(i) 
})  
// []string{"1", "2", "3"}
```

### gfn.Filter

```go
func Filter[T any](array []T, filter func(T) bool) []T
```
Filter returns a new array containing elements of the original array that satisfy the provided function.

```go
gfn.Filter([]int{1, 2, 3, 4, 5, 6}, func(i int) bool {
		return i%2 == 0
})
// []int{2, 4, 6}
```

## Map

## Math

### gfn.Abs

```go
func Abs[T Int | Float](x T) T
```

Abs returns the absolute value of x.

```go
gfn.Abs(-1)      // 1
gfn.Abs(-100.99) // 100.99
```

### gfn.Max

```go
func Max[T Int | Uint | ~float32 | ~string](array ...T) T
```

Max returns the maximum value in the array. 

> For float64 arrays, please use MaxFloat64. More details in comments of Max

```go
gfn.Max([]int16{1, 5, 9, 10}...)  // 10
gfn.Max("ab", "cd", "e")          // "e"
```

### gfn.MaxFloat64

```go
func MaxFloat64(skipNaN bool, array ...float64) float64
```

MaxFloat64 returns the maximum value in the array.

```go
gfn.MaxFloat64(false, 1.1, math.NaN())                   // NaN
gfn.MaxFloat64(true, 1.1, math.NaN(), 2.2)               // 2.2
gfn.MaxFloat64(true, []float64{math.NaN(), math.NaN(), math.NaN()}...) // NaN
```

### gfn.Min

```go
func Min[T Int | Uint | ~float32 | ~string](array ...T) T
```

Min returns the minimum value in the array.

> For float64 arrays, please use MinFloat64.

```go
gfn.Min(1.1, 2.2, 3.3)            // 1.1
gfn.Min([]int16{1, 5, 9, 10}...)  // 1
```

### gfn.MinFloat64

```go
func MinFloat64(skipNaN bool, array ...float64) float64
```

MinFloat64 returns the minimum value in the array.

```go
gfn.MinFloat64(false, math.NaN(), 1., 2., 3.)                  // NaN 
gfn.MinFloat64(true, []float64{1.1, math.Inf(-1), math.NaN()}) // math.Inf(-1)
```

### gfn.Sum

```go
func Sum[T Int | Uint | Float | ~string | Complex](array ...T) T 
```

Sum returns the sum of all values in the array.

> Be careful when using this function for float64 arrays with NaN and Inf values. `Sum([math.NaN(), 0.5])` produces `math.NaN()`. `Sum(math.Inf(1), math.Inf(-1))` produces `math.NaN()` too.

```go
gfn.Sum([]int{1, 5, 9, 10}...)  // 25
gfn.Sum(1.1, 2.2, 3.3)          // 6.6
gfn.Sum("ab", "cd", "e")        // "abcde"
```

### gfn.SumFloat64

```go
func SumFloat64(skipNaN bool, array ...float64) float64 
```

SumFloat64 returns the sum of all values in the array.

```go
gfn.SumFloat64(false, 1.1, 2.2, 3.3, math.NaN())         // NaN
gfn.SumFloat64(true, 1.1, 2.2, 3.3, math.NaN())          // 6.6
gfn.SumFloat64(true, math.Inf(1), math.Inf(-1))          // Inf(1) + Inf(-1), return NaN
```