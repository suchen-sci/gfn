# gfn
`gfn` is a Golang library that leverages generics to provide various methods, including common functional programming techniques such as `Map`, `Reduce`, and `Filter`, along with other utilities like `Contains`, `Keys`, etc.

1. No `reflect`. 
2. No third-party packages. 
3. `O(n)`.

## Why this lib?
My friend once complained to me that `Golang` is too simple, apart from the essentials, there's hardly anything else. Want to reverse an array? Not available! As a die-hard Gopher, I decided to do something, and hence this library was born. The idea of this library is very simple, it aims to port as many small utilities from other languages to `Golang` as possible. The implementation mainly refers to the methods from `Python`, `Ruby`, and `JavaScript`. I hope this library can be helpful when using `Golang`.


- [Installation](#installation)
- [Usage](#usage)
- [Type](#type)
- [Array](#array)
  - [gfn.All](#gfnall)
  - [gfn.Any](#gfnany)
  - [gfn.Contains](#gfncontains)
  - [gfn.Copy](#gfncopy)
  - [gfn.Count](#gfncount)
  - [gfn.Diff](#gfndiff)
  - [gfn.Distribution](#gfndistribution)
  - [gfn.Equal](#gfnequal)
  - [gfn.Fill](#gfnfill)
  - [gfn.GroupBy](#gfngroupby)
  - [gfn.IndexOf](#gfnindexof)
  - [gfn.IsSorted](#gfnissorted)
  - [gfn.IsSortedBy](#gfnissortedby)
  - [gfn.LastIndexOf](#gfnlastindexof)
  - [gfn.Range](#gfnrange)
  - [gfn.RangeBy](#gfnrangeby)
  - [gfn.Reverse](#gfnreverse)
  - [gfn.Sample](#gfnsample)
  - [gfn.Shuffle](#gfnshuffle)
  - [gfn.ToSet](#gfntoset)
  - [gfn.Union](#gfnunion)
  - [gfn.Uniq](#gfnuniq)
  - [gfn.Unzip](#gfnunzip)
  - [gfn.Zip](#gfnzip)
- [Functional](#functional)
  - [gfn.Filter](#gfnfilter)
  - [gfn.Map](#gfnmap)
- [Map](#map)
  - [gfn.Clear](#gfnclear)
  - [gfn.Clone](#gfnclone)
  - [gfn.Compare](#gfncompare)
  - [gfn.DeleteBy](#gfndeleteby)
  - [gfn.Invert](#gfninvert)
  - [gfn.Items](#gfnitems)
  - [gfn.Keys](#gfnkeys)
  - [gfn.Rejected](#gfnrejected)
  - [gfn.Update](#gfnupdate)
  - [gfn.Values](#gfnvalues)
- [Math](#math)
  - [gfn.Abs](#gfnabs)
  - [gfn.DivMod](#gfndivmod)
  - [gfn.Max](#gfnmax)
  - [gfn.MaxFloat64](#gfnmaxfloat64)
  - [gfn.Min](#gfnmin)
  - [gfn.MinFloat64](#gfnminfloat64)
  - [gfn.Sum](#gfnsum)



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


## Array



### gfn.All

```go
func All[T any](array []T, fn func(T) bool) bool 
```

All returns true if all elements in an array pass a given test.


```go
gfn.All([]int{1, 2, 3, 4}, func(i int) bool {
    return i > 0
}
// true
```



### gfn.Any

```go
func Any[T any](array []T, fn func(T) bool) bool 
```

Any returns true if at least one element in an array passes a given test.


```go
gfn.Any([]int{1, 2, 3, 4}, func(i int) bool {
    return i > 3
}
// true
```



### gfn.Contains

```go
func Contains[T comparable](array []T, value T) bool 
```

Contains returns true if the array contains the value.


```go
gfn.Contains([]int{1, 2, 3}, 2)             // true
gfn.Contains([]string{"a", "b", "c"}, "b")  // true
gfn.Contains([]time.Duration{time.Second}, time.Second)  // true
```



### gfn.Copy

```go
func Copy[T any](array []T) []T 
```

Copy returns a new array that is a shallow copy of the original array.


```go
gfn.Copy([]int{1, 2, 3})  // []int{1, 2, 3}

array := []int{1, 2, 3, 4, 5, 6}
gfn.Copy(array[2:])
// []int{3, 4, 5, 6}
```



### gfn.Count

```go
func Count[T comparable](array []T, value T) int 
```

Count returns the number of occurrences of a value in an array.




### gfn.Diff

```go
func Diff[T comparable](array []T, others ...[]T) []T 
```

Diff returns a new array that is a copy of the original array, removing all occurrences of any item that also appear in others. The order is preserved from the original array.


```go
gfn.Diff([]int{1, 2, 3, 4}, []int{2, 4})  // []int{1, 3}
```



### gfn.Distribution

```go
func Distribution[T comparable](array []T) map[T]int 
```

Distribution returns a map of values and their counts.


```go
gfn.Distribution([]int{1, 2, 2, 2, 2})  // map[int]int{1: 1, 2: 4}
```



### gfn.Equal

```go
func Equal[T comparable](a, b []T) bool 
```

Equal returns true if two arrays are equal.


```go
gfn.Equal([]int{1, 2, 3}, []int{1, 2, 3})                    // true
gfn.Equal([]string{"a", "c", "b"}, []string{"a", "b", "c"})  // false
```



### gfn.Fill

```go
func Fill[T any](array []T, value T) 
```

Fill sets all elements of an array to a given value.


```go
array := make([]bool, 5)
Fill(array, true)
// []bool{true, true, true, true, true}

array2 := make([]int, 5)
Fill(array2[2:], 100)
// []int{0, 0, 100, 100, 100}
```



### gfn.GroupBy

```go
func GroupBy[T any, K comparable](array []T, groupFn func(T) K) map[K][]T 
```

GroupBy generate a map of arrays by grouping the elements of an array according to a given function.


```go
array := []int{1, 2, 3, 4, 5, 6, 7, 8}
groups := GroupBy(array, func(i int) string {
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



### gfn.IndexOf

```go
func IndexOf[T comparable](array []T, value T) int 
```

IndexOf returns the index of the first occurrence of a value in an array, or -1 if not found.


```go
gfn.IndexOf([]int{1, 2, 3, 4}, 3)  // 2
gfn.IndexOf([]int{1, 2, 3, 4}, 5)  // -1
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
gfn.IsSortedBy([]int{2, 2, 1, 1, -1, -1}, func(a, b int) bool { return a >= b })
// true
```



### gfn.LastIndexOf

```go
func LastIndexOf[T comparable](array []T, value T) int 
```

LastIndexOf returns the index of the last occurrence of a value in an array, or -1 if not found.


```go
gfn.LastIndexOf([]int{3, 3, 3, 4}, 3)  // 2
gfn.LastIndexOf([]int{1, 2, 3, 4}, 5)  // -1
```



### gfn.Range

```go
func Range[T Int | Uint](start, end T) []T 
```

Range function returns a sequence of numbers, starting from start, and increments by 1, until end is reached (not included).


```go
gfn.Range(0, 7)    // []int{0, 1, 2, 3, 4, 5, 6}
gfn.Range(3, 8)    // []int{3, 4, 3, 6, 7}
gfn.Range(-10, -5) // []int{-10, -9, -8, -7, -6}
```



### gfn.RangeBy

```go
func RangeBy[T Int | Uint](start, end, step T) []T 
```

RangeBy function returns a sequence of numbers, starting from start, and increments/decrements by step, until end is reached (not included). Zero step panics.


```go
gfn.RangeBy(0, 7, 1)   // []int{0, 1, 2, 3, 4, 5, 6}
gfn.RangeBy(0, 8, 2)   // []int{0, 2, 4, 6}
gfn.RangeBy(10, 0, -2) // []int{10, 8, 6, 4, 2}
```



### gfn.Reverse

```go
func Reverse[T any](array []T) 
```

Reverse reverses an array in place.


```go
array := []int{1, 2, 3, 4}
gfn.Reverse(array)
// []int{4, 3, 2, 1}
```



### gfn.Sample

```go
func Sample[T any](array []T, n int) []T 
```

Sample returns a random sample of n elements from an array.


```go
gfn.Sample([]int{1, 2, 3, 4, 5}, 3)  // []int{3, 1, 5} or other random choices.
```



### gfn.Shuffle

```go
func Shuffle[T any](array []T) 
```

Shuffle randomizes the order of elements by using Fisher–Yates algorithm


```go
array := []int{1, 2, 3, 4}
gfn.Shuffle(array)
// array: []int{2, 1, 4, 3} or other random order
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



### gfn.Union

```go
func Union[T comparable](arrays ...[]T) []T 
```

Union returns an array with all duplicates removed from multiple arrays.


```go
gfn.Union([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
// []int{1, 2, 3, 4, 5}
```



### gfn.Uniq

```go
func Uniq[T comparable](array []T) []T 
```

Uniq returns an array with all duplicates removed.


```go
gfn.Uniq([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})  // []int{1, 2, 3, 4}
```



### gfn.Unzip

```go
func Unzip[T, U any](n int, unzipFn func(i int) (T, U)) ([]T, []U) 
```

Unzip returns two arrays built from the elements of a sequence of pairs.


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



### gfn.Zip

```go
func Zip[T, U any](a []T, b []U) []Pair[T, U] 
```

Zip returns a sequence of pairs built from the elements of two arrays.


```go
gfn.Zip([]int{1, 2, 3}, []string{"a", "b", "c"})
// []gfn.Pair[int, string]{
//     {First: 1, Second: "a"},
//     {First: 2, Second: "b"},
//     {First: 3, Second: "c"}
// }
```





## Functional



### gfn.Filter

```go
func Filter[T any](array []T, filter func(T) bool) []T 
```

Filter returns a new array containing elements of the original array that satisfy the provided function.


```go
array := []int{1, 2, 3, 4, 5, 6}
gfn.Filter(array, func(i int) bool { return i%2 == 0 })
// []int{2, 4, 6}
```



### gfn.Map

```go
func Map[T any, R any](array []T, mapper func(T) R) []R 
```

Map returns a new array with the results of calling the mapper function on each element.


```go
gfn.Map([]int{1, 2, 3}, func(i int) string {
    return strconv.Itoa(i)
})
// []string{"1", "2", "3"}

gfn.Map([]int{1, 2, 3}, func(i int) string { return i+1 })
// []int{2, 3, 4}
```





## Map



### gfn.Clear

```go
func Clear[K comparable, V any](m map[K]V) 
```

Clear removes all keys from a map.


```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
gfn.Clear(m)
// m is now an empty map
```



### gfn.Clone

```go
func Clone[K comparable, V any](m map[K]V) map[K]V 
```

Clone returns a shallow copy of a map.


```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
m2 := Clone(m)
// m2 is a copy of m
```



### gfn.Compare

```go
func Compare[K, V comparable](a map[K]V, b map[K]V) bool 
```

Compare returns true if two maps/sets are equal.


```go
map1 := map[int]struct{}{1: {}, 2: {}, 3: {}}
map2 := map[int]struct{}{1: {}, 2: {}, 3: {}}
Compare(map1, map2) // true
```



### gfn.DeleteBy

```go
func DeleteBy[K comparable, V any](m map[K]V, deleteFn func(K, V) bool) 
```

DeleteBy deletes keys from a map if the predicate function returns true.


```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
DeleteBy(m, func(k int, v string) bool {
    return k == 1 || v == "c"
})
// map[int]string{2: "b"}
```



### gfn.Invert

```go
func Invert[K, V comparable](m map[K]V) map[V]K 
```

Invert returns a map with keys and values swapped.


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



### gfn.Items

```go
func Items[K comparable, V any](m map[K]V) []Pair[K, V] 
```

Items returns a slice of pairs of keys and values.


```go
m := map[int]string{1: "a", 2: "b", 3: "c"}

gfn.Items(m)
// []Pair[int, string]{
//     {1, "a"},
//     {2, "b"},
//     {3, "c"},
// }
```



### gfn.Keys

```go
func Keys[K comparable, V any](m map[K]V) []K 
```

Keys returns the keys of a map.


```go
gfn.Keys(map[int]string{1: "a", 2: "b", 3: "c"})
// []int{1, 2, 3} or []int{3, 2, 1} or []int{2, 1, 3} etc.
```



### gfn.Rejected

```go
func Rejected[K comparable, V any](m map[K]V, keep func(K, V) bool) map[K]V 
```

Rejected returns a map with keys and values that don't pass the predicate function.


```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
rejected := Rejected(m, func(k int, v string) bool {
    return k == 1 || v == "c"
})
// map[int]string{2: "b"}
```



### gfn.Update

```go
func Update[K comparable, V any](m map[K]V, other ...map[K]V) 
```

Update updates a map with the keys and values from other maps.


```go
m1 := map[int]string{1: "a", 2: "b", 3: "c"}
m2 := map[int]string{1: "d", 2: "e"}
m3 := map[int]string{1: "f"}
Update(m1, m2, m3)
// map[int]string{1: "f", 2: "e", 3: "c"}
```



### gfn.Values

```go
func Values[K comparable, V any](m map[K]V) []V 
```

Values returns the values of a map.


```go
gfn.Values(map[int]string{1: "a", 2: "b", 3: "c"})
// []string{"a", "b", "c"} or []string{"c", "b", "a"} or []string{"b", "a", "c"} etc.
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



### gfn.DivMod

```go
func DivMod[T Int | Uint](a, b T) (T, T) 
```

DivMod returns quotient and remainder of a/b.


```go
gfn.DivMod(10, 3) // (3, 1)
```



### gfn.Max

```go
func Max[T Int | Uint | ~float32 | ~string](array ...T) T 
```

Max returns the maximum value in the array. For float64 arrays, please use MaxFloat64. NaN value in float64 arrays is not comparable to other values. Which means Max([math.NaN(), 0.5]) produces math.NaN(), but Max([0.5, math.NaN()]) produces 0.5. Since arrays with same elements but different order produce different results (inconsistent), this function does not support float64 arrays.


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

Min returns the minimum value in the array. For float64 arrays, please use MinFloat64. More details in Max.


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

Sum returns the sum of all values in the array. Be careful when using this function for float64 arrays with NaN and Inf values. Sum([math.NaN(), 0.5]) produces math.NaN(). Sum(math.Inf(1), math.Inf(-1)) produces math.NaN() too.


```go
gfn.Sum([]int{1, 5, 9, 10}...)  // 25
gfn.Sum(1.1, 2.2, 3.3)          // 6.6
gfn.Sum("ab", "cd", "e")        // "abcde"
```





