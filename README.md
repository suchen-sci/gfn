# gfn <!-- omit in toc -->
`gfn` is a Golang library that leverages generics to provide various methods, including common functional programming techniques such as `Map`, `Reduce`, and `Filter`, along with other utilities like `Contains`, `Keys`, etc.

- [Installation](#installation)
- [Usage](#usage)
- [Type](#type)
- [Array](#array)
  - [gfn.Contains](#gfncontains)
  - [gfn.Range](#gfnrange)
  - [gfn.RangeBy](#gfnrangeby)
  - [gfn.Shuffle](#gfnshuffle)
  - [gfn.Equal](#gfnequal)
  - [gfn.ToSet](#gfntoset)
- [Functional](#functional)
  - [gfn.Map](#gfnmap)
  - [gfn.Filter](#gfnfilter)
- [Map](#map)
  - [gfn.Same](#gfnsame)
- [Math](#math)
  - [gfn.Abs](#gfnabs)
  - [gfn.Max](#gfnmax)
  - [gfn.MaxFloat64](#gfnmaxfloat64)
  - [gfn.Min](#gfnmin)
  - [gfn.MinFloat64](#gfnminfloat64)
  - [gfn.Sum](#gfnsum)
  - [gfn.DivMod](#gfndivmod)
  - [gfn.IsSorted](#gfnissorted)
  - [gfn.IsSortedBy](#gfnissortedby)
  - [gfn.Distribution](#gfndistribution)


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
/*
byte: alias for uint8
rune: alias for int32
time.Duration: alias for int64
...
*/

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
gfn.Contains([]time.Duration{time.Second, 2 * time.Second}, time.Second)  // true
```

### gfn.Range

```go
func Range[T Int | Uint](start, end T) []T
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

### gfn.Shuffle

```go
func Shuffle[T any](array []T)
```

Shuffle randomizes the order of elements by using Fisher–Yates algorithm.

```go
array := []int{1, 2, 3, 4}
gfn.Shuffle(array)
// array: []int{2, 1, 4, 3} or other random order
```

### gfn.Equal

```go
func Equal[T comparable](a, b []T) bool
```

Equal returns true if two arrays are equal.

```go
gfn.Equal([]int{1, 2, 3}, []int{1, 2, 3})                // true
gfn.Equal([]string{"a", "c", "b"}, []string{"a", "b", "c"})  // false
```

### gfn.ToSet

```go
func ToSet[T comparable](array []T) map[T]struct{}
```

ToSet converts an array to a set.

```go
gfn.ToSet([]int{0, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5})
// map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}}
```

## Functional

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

### gfn.Same

```go
func Same[T, V comparable](a map[T]V, b map[T]V) bool
```

Same returns true if two maps/sets are equal.

```go
map1 := map[int]struct{}{1: {}, 2: {}, 3: {}}
map2 := map[int]struct{}{1: {}, 2: {}, 3: {}}
Same(map1, map2) // true
```

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

> For float64 arrays, please use MaxFloat64. More details in comments of Max.

```go
gfn.Max([]int16{1, 5, 9, 10}...)  // 10
gfn.Max("ab", "cd", "e")          // "e"
```

### gfn.MaxFloat64

```go
func MaxFloat64(array ...float64) float64
```

MaxFloat64 returns the maximum value in the array. NaN values are skipped.

```go
gfn.MaxFloat64(1.1, math.NaN(), 2.2)                             // 2.2
gfn.MaxFloat64([]float64{math.NaN(), math.NaN(), math.NaN()}...) // NaN
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
func MinFloat64(array ...float64) float64
```

MinFloat64 returns the minimum value in the array. NaN values are skipped.

```go
gfn.MinFloat64(1, -1, 10)                                   // -1
gfn.MinFloat64([]float64{1.1, math.Inf(-1), math.NaN()}...) // math.Inf(-1)
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

### gfn.DivMod

```go
func DivMod[T Int | Uint](a, b T) (T, T)
```

DivMod returns quotient and remainder of a/b.

```go
gfn.DivMod(10, 3)  // (3, 1)
```

### gfn.IsSorted

```go
func IsSorted[T Int | Uint | Float | ~string](array []T) bool
```

IsSorted returns true if the array is sorted in ascending order.

```go
gfn.IsSorted([]int{1, 2, 3, 4})  // true
```


### gfn.IsSortedBy

```go
func IsSortedBy[T any](array []T, order func(a1, a2 T) bool) bool
```

IsSortedBy returns true if the array is sorted in the given order. The order function should return true if a1 is ok to be placed before a2.

```go
IsSortedBy([]int{2, 2, 2, 1, 1, 1, -1, -1}, func(a, b int) bool { 
    return a >= b 
})  // true
```

### gfn.Distribution

```go
func Distribution[T comparable](array []T) map[T]int
```

Distribution returns a map of values and their counts.

```go
Distribution([]int{1, 2, 2, 2, 2})  // map[int]int{1: 1, 2: 4}
```
