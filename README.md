# gfn

[![Go Report Card](https://goreportcard.com/badge/github.com/suchen-sci/gfn?style=flat-square)](https://goreportcard.com/report/github.com/suchen-sci/gfn)
[![Coverage](https://codecov.io/gh/suchen-sci/gfn/branch/main/graph/badge.svg)](https://app.codecov.io/gh/suchen-sci/gfn/tree/main)
[![Tests](https://github.com/suchen-sci/gfn/actions/workflows/test.yml/badge.svg)](https://github.com/suchen-sci/gfn/actions/workflows/test.yml)
[![Releases](https://img.shields.io/github/release/suchen-sci/gfn/all.svg?style=flat-square)](https://github.com/suchen-sci/gfn/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/suchen-sci/gfn/blob/main/LICENSE)

`gfn` is a Golang library that leverages generics to provide various methods, including common functional programming techniques such as `Map`, `Reduce`, and `Filter`, along with other utilities like `Contains`, `Keys`, etc. The idea of this library is very simple, it aims to port as many small utilities from other languages to `Go` as possible. The implementation is highly influenced by`Python`, `Ruby`, `JavaScript` and `Lodash`.

1. No `reflect`. 
2. No third-party packages. 
3. Time complexity of `O(n)`.

- [Documentation](#documentation)
- [Contributing](#contributing)
- [License](#license)

## Documentation

- [Installation](#installation)
- [Usage](#usage)
- [Type](#type)
- [Functional](#functional)
  - [gfn.Filter](#gfnfilter)
  - [gfn.FilterKV](#gfnfilterkv)
  - [gfn.Map](#gfnmap)
  - [gfn.Reduce](#gfnreduce)
  - [gfn.ReduceKV](#gfnreducekv)
- [Math](#math)
  - [gfn.Abs](#gfnabs)
  - [gfn.DivMod](#gfndivmod)
  - [gfn.Max](#gfnmax)
  - [gfn.MaxBy](#gfnmaxby)
  - [gfn.MaxFloat64 (Deprecated)](#gfnmaxfloat64-deprecated)
  - [gfn.Mean](#gfnmean)
  - [gfn.MeanBy](#gfnmeanby)
  - [gfn.Min](#gfnmin)
  - [gfn.MinBy](#gfnminby)
  - [gfn.MinFloat64 (Deprecated)](#gfnminfloat64-deprecated)
  - [gfn.MinMax](#gfnminmax)
  - [gfn.MinMaxBy](#gfnminmaxby)
  - [gfn.MinMaxFloat64 (Deprecated)](#gfnminmaxfloat64-deprecated)
  - [gfn.Mode](#gfnmode)
  - [gfn.ModeBy](#gfnmodeby)
  - [gfn.Sum](#gfnsum)
  - [gfn.SumBy](#gfnsumby)
- [Array](#array)
  - [gfn.All](#gfnall)
  - [gfn.Any](#gfnany)
  - [gfn.Chunk](#gfnchunk)
  - [gfn.Concat](#gfnconcat)
  - [gfn.Contains](#gfncontains)
  - [gfn.Copy](#gfncopy)
  - [gfn.Count](#gfncount)
  - [gfn.CountBy](#gfncountby)
  - [gfn.Counter](#gfncounter)
  - [gfn.CounterBy](#gfncounterby)
  - [gfn.Difference](#gfndifference)
  - [gfn.DifferenceBy](#gfndifferenceby)
  - [gfn.Equal](#gfnequal)
  - [gfn.EqualBy](#gfnequalby)
  - [gfn.Fill](#gfnfill)
  - [gfn.Find](#gfnfind)
  - [gfn.FindLast](#gfnfindlast)
  - [gfn.ForEach](#gfnforeach)
  - [gfn.GroupBy](#gfngroupby)
  - [gfn.IndexOf](#gfnindexof)
  - [gfn.Intersection](#gfnintersection)
  - [gfn.IntersectionBy](#gfnintersectionby)
  - [gfn.IsSorted](#gfnissorted)
  - [gfn.IsSortedBy](#gfnissortedby)
  - [gfn.LastIndexOf](#gfnlastindexof)
  - [gfn.Range](#gfnrange)
  - [gfn.RangeBy](#gfnrangeby)
  - [gfn.Remove](#gfnremove)
  - [gfn.Repeat](#gfnrepeat)
  - [gfn.Reverse](#gfnreverse)
  - [gfn.Sample](#gfnsample)
  - [gfn.Shuffle](#gfnshuffle)
  - [gfn.ToSet](#gfntoset)
  - [gfn.Union](#gfnunion)
  - [gfn.UnionBy](#gfnunionby)
  - [gfn.Uniq](#gfnuniq)
  - [gfn.UniqBy](#gfnuniqby)
  - [gfn.Unzip](#gfnunzip)
  - [gfn.Zip](#gfnzip)
- [Map](#map)
  - [gfn.Clear](#gfnclear)
  - [gfn.Clone](#gfnclone)
  - [gfn.DeleteBy](#gfndeleteby)
  - [gfn.DifferentKeys](#gfndifferentkeys)
  - [gfn.EqualKV](#gfnequalkv)
  - [gfn.EqualKVBy](#gfnequalkvby)
  - [gfn.ForEachKV](#gfnforeachkv)
  - [gfn.GetOrDefault](#gfngetordefault)
  - [gfn.IntersectKeys](#gfnintersectkeys)
  - [gfn.Invert](#gfninvert)
  - [gfn.IsDisjoint](#gfnisdisjoint)
  - [gfn.Items](#gfnitems)
  - [gfn.Keys](#gfnkeys)
  - [gfn.Select](#gfnselect)
  - [gfn.ToKV](#gfntokv)
  - [gfn.Update](#gfnupdate)
  - [gfn.Values](#gfnvalues)



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

type Pair[T, U any] struct {
    First  T
    Second U
}
```


## Functional


### gfn.Filter
```go
func Filter[T any](array []T, filter func(T) bool) []T 
```
Filter returns a new array containing elements of the original array that satisfy the provided function.

#### Example:
```go
array := []int{1, 2, 3, 4, 5, 6}
gfn.Filter(array, func(i int) bool { return i%2 == 0 })
// []int{2, 4, 6}
```
[back to top](#gfn)


### gfn.FilterKV
```go
func FilterKV[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V 
```
FilterKV returns a new map containing elements of the original map that satisfy the provided function.

#### Example:
```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
gfn.FilterKV(m, func(k int, v string) bool {
    return k == 1 || v == "c"
})
// map[int]string{1: "a", 3: "c"}
```
[back to top](#gfn)


### gfn.Map
```go
func Map[T any, R any](array []T, mapper func(T) R) []R 
```
Map returns a new array with the results of calling the mapper function on each element. No MapKV because I don't know what to return, an array or a map? Instead, please use ForEachKV.

#### Example:
```go
gfn.Map([]int{1, 2, 3}, func(i int) string { return i+1 })
// []int{2, 3, 4}

gfn.Map([]int{1, 2, 3}, func(i int) string {
    return strconv.Itoa(i)
})
// []string{"1", "2", "3"}
```
[back to top](#gfn)


### gfn.Reduce
```go
func Reduce[T any, R any](array []T, init R, fn func(R, T) R) R 
```
Reduce executes a reducer function on each element of the array, resulting in a single output value.

#### Example:
```go
gfn.Reduce([]int{1, 2, 3}, 0, func(a, b int) int {
    return a + b
})
// 6
```
[back to top](#gfn)


### gfn.ReduceKV
```go
func ReduceKV[K comparable, V any, R any](m map[K]V, init R, fn func(R, K, V) R) R 
```
ReduceKV executes a reducer function on each element of the map, resulting in a single output value.

#### Example:
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
total := gfn.ReduceKV(m, 0, func(value int, k string, v int) int {
    return value + v
})
// 6
```
[back to top](#gfn)




## Math


### gfn.Abs
```go
func Abs[T Int | Float](x T) T 
```
Abs returns the absolute value of x.

#### Example:
```go
gfn.Abs(-1)      // 1
gfn.Abs(-100.99) // 100.99
```
[back to top](#gfn)


### gfn.DivMod
```go
func DivMod[T Int | Uint](a, b T) (T, T) 
```
DivMod returns quotient and remainder of a/b.

#### Example:
```go
gfn.DivMod(10, 3) // (3, 1)
```
[back to top](#gfn)


### gfn.Max
```go
func Max[T Int | Uint | Float | ~string](array ...T) T 
```
Max returns the maximum value in the array. For float64 arrays, NaN values are skipped.

#### Example:
```go
gfn.Max([]int16{1, 5, 9, 10}...)  // 10
gfn.Max("ab", "cd", "e")          // "e"

gfn.Max(1.1, math.NaN(), 2.2)                             // 2.2
gfn.Max([]float64{math.NaN(), math.NaN(), math.NaN()}...) // NaN
```
[back to top](#gfn)


### gfn.MaxBy
```go
func MaxBy[T any, U Int | Uint | Float | ~string](array []T, fn func(T) U) T 
```
MaxBy returns the maximum value in the array, using the given function to transform values.

#### Example:
```go
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
```
[back to top](#gfn)


### gfn.MaxFloat64 (Deprecated)
```go
func MaxFloat64(array ...float64) float64 
```
Deprecated: MaxFloat64 returns the maximum value in the array. Use Max instead.

#### Example:
```go
gfn.MaxFloat64(1.1, math.NaN(), 2.2)                             // 2.2
gfn.MaxFloat64([]float64{math.NaN(), math.NaN(), math.NaN()}...) // NaN
```
[back to top](#gfn)


### gfn.Mean
```go
func Mean[T Int | Uint | Float](array ...T) float64 
```
Mean returns the mean of all values in the array.

#### Example:
```go
gfn.Mean(1, 2, 3)               // 2.0
gfn.Mean([]int{1, 2, 3, 4}...)  // 2.5
```
[back to top](#gfn)


### gfn.MeanBy
```go
func MeanBy[T any, U Int | Uint | Float](array []T, fn func(T) U) float64 
```
MeanBy returns the mean of all values in the array after applying fn to each value.

#### Example:
```go
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
```
[back to top](#gfn)


### gfn.Min
```go
func Min[T Int | Uint | Float | ~string](array ...T) T 
```
Min returns the minimum value in the array. For float64 arrays, NaN values are skipped.

#### Example:
```go
gfn.Min(1.1, 2.2, 3.3)            // 1.1
gfn.Min([]int16{1, 5, 9, 10}...)  // 1

gfn.Min(1, -1, 10)                                   // -1
gfn.Min([]float64{1.1, math.Inf(-1), math.NaN()}...) // math.Inf(-1)
```
[back to top](#gfn)


### gfn.MinBy
```go
func MinBy[T any, U Int | Uint | Float | ~string](array []T, fn func(T) U) T 
```
MinBy returns the minimum value in the array, using the given function to transform values.

#### Example:
```go
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
```
[back to top](#gfn)


### gfn.MinFloat64 (Deprecated)
```go
func MinFloat64(array ...float64) float64 
```
Deprecated: MinFloat64 returns the minimum value in the array. Use Min instead.

#### Example:
```go
gfn.MinFloat64(1, -1, 10)                                   // -1
gfn.MinFloat64([]float64{1.1, math.Inf(-1), math.NaN()}...) // math.Inf(-1)
```
[back to top](#gfn)


### gfn.MinMax
```go
func MinMax[T Int | Uint | Float | ~string](array ...T) (T, T) 
```
MinMax returns the minimum and maximum value in the array. For float64 arrays, please use MinMaxFloat64.

#### Example:
```go
gfn.MinMax(1, 5, 9, 10)  // 1, 10

gfn.MinMax(math.NaN(), 1.85, 2.2) // 1.85, 2.2
gfn.MinMax(math.NaN(), math.NaN(), math.NaN()) // NaN, NaN
```
[back to top](#gfn)


### gfn.MinMaxBy
```go
func MinMaxBy[T any, U Int | Uint | Float | ~string](array []T, fn func(T) U) (T, T) 
```
MinMaxBy returns the minimum and maximum value in the array, using the given function to transform values.

#### Example:
```go
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
```
[back to top](#gfn)


### gfn.MinMaxFloat64 (Deprecated)
```go
func MinMaxFloat64(array ...float64) (float64, float64) 
```
Deprecated: MinMaxFloat64 returns the minimum and maximum value in the array. Use MinMax instead.

#### Example:
```go
gfn.MinMaxFloat64(math.NaN(), 1.85, 2.2) // 1.85, 2.2
gfn.MinMaxFloat64(math.NaN(), math.NaN(), math.NaN()) // NaN, NaN
```
[back to top](#gfn)


### gfn.Mode
```go
func Mode[T comparable](array []T) T 
```
Mode returns the most frequent value in the array.

#### Example:
```go
gfn.Mode([]int{1, 1, 5, 5, 5, 2, 2})) // 5
```
[back to top](#gfn)


### gfn.ModeBy
```go
func ModeBy[T any, U comparable](array []T, fn func(T) U) T 
```
ModeBy returns the most frequent value in the array, using the given function to transform values.

#### Example:
```go
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
```
[back to top](#gfn)


### gfn.Sum
```go
func Sum[T Int | Uint | Float | ~string | Complex](array ...T) T 
```
Sum returns the sum of all values in the array. Be careful when using this function for float64 arrays with NaN and Inf values. Sum([math.NaN(), 0.5]) produces math.NaN(). Sum(math.Inf(1), math.Inf(-1)) produces math.NaN() too.

#### Example:
```go
gfn.Sum([]int{1, 5, 9, 10}...)  // 25
gfn.Sum(1.1, 2.2, 3.3)          // 6.6
gfn.Sum("ab", "cd", "e")        // "abcde"
```
[back to top](#gfn)


### gfn.SumBy
```go
func SumBy[T any, U Int | Uint | Float | ~string](array []T, fn func(T) U) U 
```
SumBy returns the sum of all values in the array after applying fn to each value.

#### Example:
```go
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
```
[back to top](#gfn)




## Array


### gfn.All
```go
func All[T any](array []T, fn func(T) bool) bool 
```
All returns true if all elements in an array pass a given test.

#### Example:
```go
gfn.All([]int{1, 2, 3, 4}, func(i int) bool {
    return i > 0
}
// true
```
[back to top](#gfn)


### gfn.Any
```go
func Any[T any](array []T, fn func(T) bool) bool 
```
Any returns true if at least one element in an array passes a given test.

#### Example:
```go
gfn.Any([]int{1, 2, 3, 4}, func(i int) bool {
    return i > 3
}
// true
```
[back to top](#gfn)


### gfn.Chunk
```go
func Chunk[T any](array []T, size int) [][]T 
```
Chunk splits an array into chunks of given size.

#### Example:
```go
gfn.Chunk([]int{1, 2, 3, 4, 5}, 2)  // [][]int{{1, 2}, {3, 4}, {5}}
```
[back to top](#gfn)


### gfn.Concat
```go
func Concat[T any](arrays ...[]T) []T 
```
Concat returns a new array that is the result of joining two or more arrays.

#### Example:
```go
gfn.Concat([]int{1, 2}, []int{3, 4})  // []int{1, 2, 3, 4}
```
[back to top](#gfn)


### gfn.Contains
```go
func Contains[T comparable](array []T, value T) bool 
```
Contains returns true if the array contains the value.

#### Example:
```go
gfn.Contains([]int{1, 2, 3}, 2)             // true
gfn.Contains([]string{"a", "b", "c"}, "b")  // true
gfn.Contains([]time.Duration{time.Second}, time.Second)  // true
```
[back to top](#gfn)


### gfn.Copy
```go
func Copy[T any](array []T) []T 
```
Copy returns a new array that is a shallow copy of the original array.

#### Example:
```go
gfn.Copy([]int{1, 2, 3})  // []int{1, 2, 3}

array := []int{1, 2, 3, 4, 5, 6}
gfn.Copy(array[2:])
// []int{3, 4, 5, 6}
```
[back to top](#gfn)


### gfn.Count
```go
func Count[T comparable](array []T, value T) int 
```
Count returns the number of occurrences of a value in an array.

#### Example:
```go
gfn.Count([]int{1, 2, 2, 2, 5, 6}, 2)  // 3
```
[back to top](#gfn)


### gfn.CountBy
```go
func CountBy[T any](array []T, fn func(T) bool) int 
```
CountBy returns the number of elements in an array that satisfy a predicate.

#### Example:
```go
type Employee struct {
    name       string
    department string
}
employees := []Employee{
    {"Alice", "Accounting"},
    {"Cindy", "Engineering"},
    {"Dave", "Engineering"},
    {"Eve", "Engineering"},
}
gfn.CountBy(employees, func(e Employee) bool {
    return e.department == "Engineering"
})  // 3
```
[back to top](#gfn)


### gfn.Counter
```go
func Counter[T comparable](array []T) map[T]int 
```
Counter returns a map of values and their counts.

#### Example:
```go
gfn.Counter([]int{1, 2, 2, 2, 2})  // map[int]int{1: 1, 2: 4}
```
[back to top](#gfn)


### gfn.CounterBy
```go
func CounterBy[T any, U comparable](array []T, fn func(T) U) map[U]int 
```
CounterBy returns a map of values and their counts. The values are calculated by the given function.

#### Example:
```go
type Employee struct {
    name       string
    department string
}
employees := []Employee{
    {"Alice", "Accounting"},
    {"Dave", "Engineering"},
    {"Eve", "Engineering"},
}
gfn.CounterBy(employees, func(e Employee) string {
    return e.department
})  // map[string]int{"Accounting": 1, "Engineering": 2}
```
[back to top](#gfn)


### gfn.Difference
```go
func Difference[T comparable](array []T, others ...[]T) []T 
```
Difference returns a new array that is a copy of the original array, removing all occurrences of any item that also appear in others. The order is preserved from the original array.

#### Example:
```go
gfn.Difference([]int{1, 2, 3, 4}, []int{2, 4})  // []int{1, 3}
```
[back to top](#gfn)


### gfn.DifferenceBy
```go
func DifferenceBy[T any, U comparable](fn func(T) U, array []T, others ...[]T) []T 
```
DifferenceBy returns a new array that is a copy of the original array, removing all occurrences of any item that also appear in others. The occurrences are determined by applying a function to each element.

#### Example:
```go
type Data struct {
    value int
}
data1 := []Data{{1}, {3}, {2}, {4}, {5}, {2}}
data2 := []Data{{3}, {4}, {5}}
gfn.DifferenceBy(func(d Data) int { return d.value }, data1, data2)
// []Data{{1}, {2}, {2}}
```
[back to top](#gfn)


### gfn.Equal
```go
func Equal[T comparable](a, b []T) bool 
```
Equal returns true if two arrays are equal. Two arrays are considered equal if both are nil, or if their lengths are equal and their elements are equal. Elements are compared using == operator.

#### Example:
```go
gfn.Equal([]int{1, 2, 3}, []int{1, 2, 3})                    // true
gfn.Equal([]string{"a", "c", "b"}, []string{"a", "b", "c"})  // false
```
[back to top](#gfn)


### gfn.EqualBy
```go
func EqualBy[T1, T2 any](a []T1, b []T2, fn func(T1, T2) bool) bool 
```
EqualBy returns true if two arrays are equal by comparing their elements using the given function.

#### Example:
```go
a := []int{1, 2, 3, 4, 5}
b := []rune{'a', 'b', 'c', 'd', 'e'}
gfn.EqualBy(a, b, func(aa int, bb rune) bool {
    return (aa - 1) == int(bb-'a')
}) // true
```
[back to top](#gfn)


### gfn.Fill
```go
func Fill[T any](array []T, value T) 
```
Fill sets all elements of an array to a given value. You can control the start and end index by using the slice.

#### Example:
```go
array := make([]bool, 5)
gfn.Fill(array, true)
// []bool{true, true, true, true, true}

// you can control the start and end index by using the slice
array2 := make([]int, 5)
gfn.Fill(array2[2:], 100)
// []int{0, 0, 100, 100, 100}
```
[back to top](#gfn)


### gfn.Find
```go
func Find[T any](array []T, fn func(T) bool) (T, int) 
```
Find returns the first element in an array that passes a given test and corresponding index. Index of -1 is returned if no element passes the test.

#### Example:
```go
value, index := gfn.Find([]string{"a", "ab", "abc"}, func(s string) bool {
    return len(s) > 1
})
// "ab", 1
```
[back to top](#gfn)


### gfn.FindLast
```go
func FindLast[T any](array []T, fn func(T) bool) (T, int) 
```
FindLast returns the last element in an array that passes a given test and corresponding index. Index of -1 is returned if no element passes the test.

#### Example:
```go
value, index := gfn.FindLast([]string{"a", "ab", "abc"}, func(s string) bool {
    return len(s) > 1
})
// "abc", 2
```
[back to top](#gfn)


### gfn.ForEach
```go
func ForEach[T any](array []T, fn func(value T)) 
```
ForEach executes a provided function once for each array element.

#### Example:
```go
sum := 0
gfn.ForEach([]int{1, 2, 3}, func(i int) {
    sum += i
})
// sum == 6
```
[back to top](#gfn)


### gfn.GroupBy
```go
func GroupBy[T any, K comparable](array []T, groupFn func(T) K) map[K][]T 
```
GroupBy generate a map of arrays by grouping the elements of an array according to a given function.

#### Example:
```go
array := []int{1, 2, 3, 4, 5, 6, 7, 8}
groups := gfn.GroupBy(array, func(i int) string {
    if i%2 == 0 {
        return "even"
    }
    return "odd"
})
// map[string][]int{
//     "even": []int{2, 4, 6, 8},
//     "odd":  []int{1, 3, 5, 7},
// }
```
[back to top](#gfn)


### gfn.IndexOf
```go
func IndexOf[T comparable](array []T, value T) int 
```
IndexOf returns the index of the first occurrence of a value in an array, or -1 if not found.

#### Example:
```go
gfn.IndexOf([]int{1, 2, 3, 4}, 3)  // 2
gfn.IndexOf([]int{1, 2, 3, 4}, 5)  // -1
```
[back to top](#gfn)


### gfn.Intersection
```go
func Intersection[T comparable](arrays ...[]T) []T 
```
Intersection returns a new array that is the intersection of two or more arrays.

#### Example:
```go
arr1 := []int{1, 2, 3, 4, 5}
arr2 := []int{2, 3, 4, 5, 6}
arr3 := []int{5, 4, 3, 2}
arr4 := []int{2, 3}
gfn.Intersection(arr1, arr2, arr3, arr4)  // []int{2, 3}
```
[back to top](#gfn)


### gfn.IntersectionBy
```go
func IntersectionBy[T any, U comparable](fn func(T) U, arrays ...[]T) []T 
```
IntersectionBy returns a new array that is the intersection of two or more arrays, where intersection is determined by a given function.

#### Example:
```go
type Data struct {
    value int
}
data1 := []Data{{1}, {3}, {2}, {4}, {5}}
data2 := []Data{{2}, {3}}
gfn.IntersectionBy(func(d Data) int { return d.value }, data1, data2)
// []Data{{3}, {2}}
```
[back to top](#gfn)


### gfn.IsSorted
```go
func IsSorted[T Int | Uint | Float | ~string](array []T) bool 
```
IsSorted returns true if the array is sorted in ascending order.

#### Example:
```go
gfn.IsSorted([]int{1, 2, 3, 4})  // true
```
[back to top](#gfn)


### gfn.IsSortedBy
```go
func IsSortedBy[T any](array []T, order func(a1, a2 T) bool) bool 
```
IsSortedBy returns true if the array is sorted in the given order. The order function should return true if a1 is ok to be placed before a2.

#### Example:
```go
gfn.IsSortedBy([]int{2, 2, 1, 1, -1, -1}, func(a, b int) bool { return a >= b })
// true
```
[back to top](#gfn)


### gfn.LastIndexOf
```go
func LastIndexOf[T comparable](array []T, value T) int 
```
LastIndexOf returns the index of the last occurrence of a value in an array, or -1 if not found.

#### Example:
```go
gfn.LastIndexOf([]int{3, 3, 3, 4}, 3)  // 2
gfn.LastIndexOf([]int{1, 2, 3, 4}, 5)  // -1
```
[back to top](#gfn)


### gfn.Range
```go
func Range[T Int | Uint](start, end T) []T 
```
Range function returns a sequence of numbers, starting from start, and increments by 1, until end is reached (not included).

#### Example:
```go
gfn.Range(0, 7)    // []int{0, 1, 2, 3, 4, 5, 6}
gfn.Range(3, 8)    // []int{3, 4, 3, 6, 7}
gfn.Range(-10, -5) // []int{-10, -9, -8, -7, -6}
```
[back to top](#gfn)


### gfn.RangeBy
```go
func RangeBy[T Int | Uint](start, end, step T) []T 
```
RangeBy function returns a sequence of numbers, starting from start, and increments/decrements by step, until end is reached (not included). Zero step panics.

#### Example:
```go
gfn.RangeBy(0, 7, 1)   // []int{0, 1, 2, 3, 4, 5, 6}
gfn.RangeBy(0, 8, 2)   // []int{0, 2, 4, 6}
gfn.RangeBy(10, 0, -2) // []int{10, 8, 6, 4, 2}
```
[back to top](#gfn)


### gfn.Remove
```go
func Remove[T comparable](array []T, values ...T) []T 
```
Remove removes all elements from an array that equal to given values.

#### Example:
```go
gfn.Remove([]int{1, 2, 3, 4, 2, 3, 2, 3}, 2, 3)  // []int{1, 4}
```
[back to top](#gfn)


### gfn.Repeat
```go
func Repeat[T any](array []T, repeat int) []T 
```
Repeat returns a new array that is the result of repeating an array a given number of times.

#### Example:
```go
gfn.Repeat([]int{1, 2, 3}, 3)  // []int{1, 2, 3, 1, 2, 3, 1, 2, 3}
```
[back to top](#gfn)


### gfn.Reverse
```go
func Reverse[T any](array []T) 
```
Reverse reverses an array in place.

#### Example:
```go
array := []int{1, 2, 3, 4}
gfn.Reverse(array)
// []int{4, 3, 2, 1}
```
[back to top](#gfn)


### gfn.Sample
```go
func Sample[T any](array []T, n int) []T 
```
Sample returns a random sample of n elements from an array. Every position in the array are at most selected once. n should be less or equal to len(array).

#### Example:
```go
gfn.Sample([]int{1, 2, 3, 4, 5}, 3)  // []int{3, 1, 5} or other random choices.
```
[back to top](#gfn)


### gfn.Shuffle
```go
func Shuffle[T any](array []T) 
```
Shuffle randomizes the order of elements by using Fisher–Yates algorithm

#### Example:
```go
array := []int{1, 2, 3, 4}
gfn.Shuffle(array)
// array: []int{2, 1, 4, 3} or other random order
```
[back to top](#gfn)


### gfn.ToSet
```go
func ToSet[T comparable](array []T) map[T]struct{} 
```
ToSet converts an array to a set.

#### Example:
```go
gfn.ToSet([]int{0, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5})
// map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}}
```
[back to top](#gfn)


### gfn.Union
```go
func Union[T comparable](arrays ...[]T) []T 
```
Union returns an array with all duplicates removed from multiple arrays.

#### Example:
```go
gfn.Union([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
// []int{1, 2, 3, 4, 5}
```
[back to top](#gfn)


### gfn.UnionBy
```go
func UnionBy[T any, U comparable](fn func(T) U, arrays ...[]T) []T 
```
UnionBy returns an array with all duplicates removed from multiple arrays by applying a function to each element.

#### Example:
```go
type Employee struct {
    name       string
    department string
}
group1 := []Employee{
    {"Alice", "Accounting"},
    {"Bob", "Accounting"},
    {"Cindy", "Engineering"},
}
group2 := []Employee{
    {"Alice", "Accounting"},
    {"Cindy", "Engineering"},
    {"Dave", "Engineering"},
    {"Eve", "Engineering"},
}
gfn.UnionBy(func(e Employee) string { return e.name }, group1, group2)
// []Employee{
//     {"Alice", "Accounting"},
//     {"Bob", "Accounting"},
//     {"Cindy", "Engineering"},
//     {"Dave", "Engineering"},
//     {"Eve", "Engineering"},
// }
```
[back to top](#gfn)


### gfn.Uniq
```go
func Uniq[T comparable](array []T) []T 
```
Uniq returns an array with all duplicates removed.

#### Example:
```go
gfn.Uniq([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})  // []int{1, 2, 3, 4}
```
[back to top](#gfn)


### gfn.UniqBy
```go
func UniqBy[T any, U comparable](array []T, fn func(T) U) []T 
```
UniqBy returns an array with all duplicates removed by applying a function to each element.

#### Example:
```go
type Employee struct {
    name       string
    department string
}
employees := []Employee{
    {"Alice", "Accounting"},
    {"Bob", "Accounting"},
    {"Cindy", "Engineering"},
    {"Dave", "Engineering"},
}
gfn.UniqBy(employees, func(e Employee) string {
    return e.department
})
// []Employee{{"Alice", "Accounting"}, {"Cindy", "Engineering"}}
```
[back to top](#gfn)


### gfn.Unzip
```go
func Unzip[T, U any](n int, unzipFn func(i int) (T, U)) ([]T, []U) 
```
Unzip returns two arrays built from the elements of a sequence of pairs.

#### Example:
```go
pairs := []gfn.Pair[int, string]{
    {First: 1, Second: "a"},
    {First: 2, Second: "b"},
    {First: 3, Second: "c"},
}
gfn.Unzip(len(pairs), func(i int) (int, string) {
    return pairs[i].First, pairs[i].Second
})
// ([]int{1, 2, 3}, []string{"a", "b", "c"})
```
[back to top](#gfn)


### gfn.Zip
```go
func Zip[T, U any](a []T, b []U) []Pair[T, U] 
```
Zip returns a sequence of pairs built from the elements of two arrays.

#### Example:
```go
gfn.Zip([]int{1, 2, 3}, []string{"a", "b", "c"})
// []gfn.Pair[int, string]{
//     {First: 1, Second: "a"},
//     {First: 2, Second: "b"},
//     {First: 3, Second: "c"}
// }
```
[back to top](#gfn)




## Map


### gfn.Clear
```go
func Clear[K comparable, V any](m map[K]V) 
```
Clear removes all keys from a map.

#### Example:
```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
gfn.Clear(m)
// m is now an empty map
```
[back to top](#gfn)


### gfn.Clone
```go
func Clone[K comparable, V any](m map[K]V) map[K]V 
```
Clone returns a shallow copy of a map.

#### Example:
```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
m2 := gfn.Clone(m)
// m2 is a copy of m
```
[back to top](#gfn)


### gfn.DeleteBy
```go
func DeleteBy[K comparable, V any](m map[K]V, deleteFn func(K, V) bool) 
```
DeleteBy deletes keys from a map if the predicate function returns true.

#### Example:
```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
gfn.DeleteBy(m, func(k int, v string) bool {
    return k == 1 || v == "c"
})
// map[int]string{2: "b"}
```
[back to top](#gfn)


### gfn.DifferentKeys
```go
func DifferentKeys[K comparable, V any](ms ...map[K]V) []K 
```
DifferentKeys returns a slice of keys that are in the first map but not in the others, only keys in the map are considered, not values. It usually used to find the difference between two or more sets.

#### Example:
```go
m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
m2 := map[int]string{1: "a", 2: "b"}
m3 := map[int]string{2: "b", 3: "c"}
gfn.DifferentKeys(m1, m2, m3)  // []int{4}
```
[back to top](#gfn)


### gfn.EqualKV
```go
func EqualKV[K, V comparable](a map[K]V, b map[K]V) bool 
```
EqualKV returns true if two maps/sets are equal.

#### Example:
```go
map1 := map[int]struct{}{1: {}, 2: {}, 3: {}}
map2 := map[int]struct{}{1: {}, 2: {}, 3: {}}
gfn.EqualKV(map1, map2) // true
```
[back to top](#gfn)


### gfn.EqualKVBy
```go
func EqualKVBy[K comparable, V1, V2 any](a map[K]V1, b map[K]V2, fn func(K, V1, V2) bool) bool 
```
EqualKVBy returns true if two maps/sets are equal by a custom function.

#### Example:
```go
m1 := map[int]string{1: "a", 2: "b", 3: "c"}
m2 := map[int]string{1: "e", 2: "f", 3: "g"}
gfn.EqualKVBy(m1, m2, func(k int, a, b string) bool {
    return len(a) == len(b)
}) // true
```
[back to top](#gfn)


### gfn.ForEachKV
```go
func ForEachKV[K comparable, V any](m map[K]V, fn func(K, V)) 
```
ForEachKV iterates over a map and calls a function for each key/value pair.

#### Example:
```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
array := make([]int, 0, len(m))
gfn.ForEachKV(m, func(k int, v string) {
    array = append(array, k)
}
// array is []int{1, 2, 3} or other order

m := map[int]string{1: "a", 2: "b", 3: "c"}
invert := map[string]int{}
gfn.ForEachKV(m, func(k int, v string) {
    invert[v] = k
}
// invert is map[string]int{"a": 1, "b": 2, "c": 3}
```
[back to top](#gfn)


### gfn.GetOrDefault
```go
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V 
```
GetOrDefault returns the value for a key if it exists, otherwise it returns the default value.

#### Example:
```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
gfn.GetOrDefault(m, 1, "d")  // "a"
gfn.GetOrDefault(m, 4, "d")  // "d"
```
[back to top](#gfn)


### gfn.IntersectKeys
```go
func IntersectKeys[K comparable, V any](ms ...map[K]V) []K 
```
IntersectKeys returns a slice of keys that are in all maps. It usually used to find the intersection of two or more sets.

#### Example:
```go
m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
m2 := map[int]string{1: "a", 2: "b"}
m3 := map[int]string{2: "b", 3: "c", 4: "d"}
gfn.IntersectKeys(m1, m2, m3)  // []int{2}
```
[back to top](#gfn)


### gfn.Invert
```go
func Invert[K, V comparable](m map[K]V) map[V]K 
```
Invert returns a map with keys and values swapped.

#### Example:
```go
m := map[string]string{
    "Array": "array.go",
    "Map":   "map.go",
    "Set":   "set.go",
    "Math":  "math.go",
}

gfn.Invert(m)
// map[string]string{
//     "array.go": "Array",
//     "map.go":   "Map",
//     "set.go":   "Set",
//     "math.go":  "Math",
// }
```
[back to top](#gfn)


### gfn.IsDisjoint
```go
func IsDisjoint[K comparable, V1 any, V2 any](m1 map[K]V1, m2 map[K]V2) bool 
```
IsDisjoint returns true if the maps have no keys in common. It usually used to check if two sets are disjoint.

#### Example:
```go
m1 := map[int]string{1: "a", 2: "b", 3: "c"}
m2 := map[int]int{4: 4, 5: 5, 6: 6}
IsDisjoint(m1, m2)  // true

m3 := map[int]struct{}{1: {}, 2: {}, 3: {}}
m4 := map[int]struct{}{4: {}, 5: {}, 6: {}}
gfn.IsDisjoint(m3, m4)  // true
```
[back to top](#gfn)


### gfn.Items
```go
func Items[K comparable, V any](m map[K]V) []Pair[K, V] 
```
Items returns a slice of pairs of keys and values.

#### Example:
```go
m := map[int]string{1: "a", 2: "b", 3: "c"}

gfn.Items(m)
// []gfn.Pair[int, string]{
//     {1, "a"},
//     {2, "b"},
//     {3, "c"},
// }
```
[back to top](#gfn)


### gfn.Keys
```go
func Keys[K comparable, V any](m map[K]V) []K 
```
Keys returns the keys of a map.

#### Example:
```go
gfn.Keys(map[int]string{1: "a", 2: "b", 3: "c"})
// []int{1, 2, 3} or []int{3, 2, 1} or []int{2, 1, 3} etc.
```
[back to top](#gfn)


### gfn.Select
```go
func Select[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V 
```
Select returns a map with keys and values that satisfy the predicate function.

#### Example:
```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
gfn.Select(m, func(k int, v string) bool {
    return k == 1 || v == "c"
})
// map[int]string{1: "a", 3: "c"}
```
[back to top](#gfn)


### gfn.ToKV
```go
func ToKV[K comparable, V any](n int, fn func(int) (K, V)) map[K]V 
```
ToKV converts a slice to a map using a function to generate the keys and values.

#### Example:
```go
gfn.ToKV(3, func(i int) (int, string) {
    return i, strconv.Itoa(i)
})
// map[int]string{0: "0", 1: "1", 2: "2"}
```
[back to top](#gfn)


### gfn.Update
```go
func Update[K comparable, V any](m map[K]V, other ...map[K]V) 
```
Update updates a map with the keys and values from other maps.

#### Example:
```go
// use Update to do union of maps
m1 := map[int]string{1: "a", 2: "b", 3: "c"}
m2 := map[int]string{4: "d", 5: "e"}
union := map[int]string{}
gfn.Update(union, m1, m2)
// union: map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}

m1 := map[int]string{1: "a", 2: "b", 3: "c"}
m2 := map[int]string{1: "d", 2: "e"}
m3 := map[int]string{1: "f"}
gfn.Update(m1, m2, m3)
// map[int]string{1: "f", 2: "e", 3: "c"}
```
[back to top](#gfn)


### gfn.Values
```go
func Values[K comparable, V any](m map[K]V) []V 
```
Values returns the values of a map.

#### Example:
```go
gfn.Values(map[int]string{1: "a", 2: "b", 3: "c"})
// []string{"a", "b", "c"} or []string{"c", "b", "a"} or []string{"b", "a", "c"} etc.
```
[back to top](#gfn)





## Contributing

Format:
```go
/* @example MyFunc
// your examples here
MyFunc(...)
// output: ...
*/

// MyFunc is ...
func MyFunc(args ...) (return values...) {
    // your code here.
}
```

then run following command to update `README.md` (generated from `README.tmpl.md`).

```bash
make doc
```

[back to top](#gfn)

## License
`gfn` is under the MIT license. See the [LICENSE](https://github.com/suchen-sci/gfn/blob/main/LICENSE) file for details.


[back to top](#gfn)
