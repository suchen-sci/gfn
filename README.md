# gfn <!-- omit in toc -->
`gfn` is a Golang library that leverages generics to provide various methods, including common functional programming techniques such as `Map`, `Reduce`, and `Filter`, along with other utilities like `Contains`, `Keys`, etc.

- [Installation](#installation)
- [Usage](#usage)
- [Array](#array)
  - [gfn.Contains](#gfncontains)
  - [gfn.Map](#gfnmap)
  - [gfn.Filter](#gfnfilter)
- [Map](#map)
- [Math](#math)
  - [gfn.Max](#gfnmax)
  - [gfn.MaxNotNaN](#gfnmaxnotnan)
  - [gfn.Min](#gfnmin)
  - [gfn.MinNotNaN](#gfnminnotnan)
  - [gfn.Sum](#gfnsum)
  - [gfn.SumNotNaN](#gfnsumnotnan)

## Installation
```
go get github.com/suchen-sci/gfn
```

## Usage 
```
import "github.com/suchen-sci/gfn"
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

### gfn.Max

```go
func Max[T Int | Uint | Float | ~string](array ...T) T
```

Max returns the maximum value in the array.

> Be careful when using this function for float64 arrays with NaN values. NaN values is not comparable to other float64. NaN > x is false. NaN < x is false. Which means `Max([math.NaN(), 0.5])` is `math.NaN()`, but `Max([0.5, math.NaN()])` is `0.5`.

```go
gfn.Max([]int16{1, 5, 9, 10}...)  // 10
gfn.Max("ab", "cd", "e")          // "e"
```

### gfn.MaxNotNaN

```go
func MaxNotNaN(array ...float64) float64
```

MaxNotNaN returns the maximum not NaN value in the array.

```go
MaxNotNaN(1., 2., 3., math.NaN())              // 3.
MaxNotNaN(math.NaN(), math.NaN(), math.NaN())  // all NaN, return NaN
```

### gfn.Min

```go
func Min[T Int | Uint | Float | ~string](array ...T) T
```

Min returns the minimum value in the array.
> Be careful when using this function for float64 arrays with NaN values. NaN values is not comparable to other float64. NaN > x is false. NaN < x is false. Which means `Min([math.NaN(), 0.5])` is `math.NaN()`, but `Min([0.5, math.NaN()])` is `0.5`.

```go
gfn.Min(1.1, 2.2, 3.3)            // 1.1
gfn.Min([]int16{1, 5, 9, 10}...)  // 1
```

### gfn.MinNotNaN

```go
func MinNotNaN(array ...float64) float64
```

MinNotNaN returns the minimum not NaN value in the array.

```go
gfn.MinNotNaN(math.NaN(), 1., 2., 3.)             // 1. 
gfn.MinNotNaN(math.NaN(), math.NaN(), math.NaN()) // all NaN, return NaN
```

### gfn.Sum

```go
func Sum[T Int | Uint | Float | ~string | Complex](array ...T) T 
```

Sum returns the sum of all values in the array.
> Be careful when using this function for float64 arrays with NaN and Inf values. NaN values is not comparable to other float64. NaN + x is NaN. Which means `Sum([math.NaN(), 0.5])` is `math.NaN()`. `Sum(math.Inf(1), math.Inf(-1))` is `math.NaN()` too.

```go
gfn.Sum([]int{1, 5, 9, 10}...)  // 25
gfn.Sum(1.1, 2.2, 3.3)          // 6.6
gfn.Sum("ab", "cd", "e")        // "abcde"
```

### gfn.SumNotNaN

```go
func SumNotNaN(array ...float64) float64
```

SumNotNaN returns the sum not NaN values in the array.

```go
gfn.SumNotNaN(1.1, 2.2, 3.3, math.NaN())          // 6.6
gfn.SumNotNaN(math.NaN(), math.NaN(), math.NaN()) // all NaN, return NaN
gfn.SumNotNaN(math.Inf(1), math.Inf(-1))          // Inf(1) + Inf(-1), return NaN
```