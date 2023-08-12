// gfn is a Golang library that leverages generics to provide various methods.
package gfn

import "math/rand"

// Contains returns true if the array contains the value.
// @example
// gfn.Contains([]int{1, 2, 3}, 2)             // true
// gfn.Contains([]string{"a", "b", "c"}, "b")  // true
// gfn.Contains([]time.Duration{time.Second}, time.Second)  // true
func Contains[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// Range function returns a sequence of numbers, starting from start,
// and increments by 1, until end is reached (not included).
// @example
// gfn.Range(0, 7)    // []int{0, 1, 2, 3, 4, 5, 6}
// gfn.Range(3, 8)    // []int{3, 4, 3, 6, 7}
// gfn.Range(-10, -5) // []int{-10, -9, -8, -7, -6}
func Range[T Int | Uint](start, end T) []T {
	if start >= end {
		return []T{}
	}

	res := make([]T, end-start)
	for i := 0; i < len(res); i++ {
		res[i] = start + T(i)
	}
	return res
}

// RangeBy function returns a sequence of numbers, starting from start,
// and increments/decrements by step, until end is reached (not included).
// Zero step panics.
// @example
// gfn.RangeBy(0, 7, 1)   // []int{0, 1, 2, 3, 4, 5, 6}
// gfn.RangeBy(0, 8, 2)   // []int{0, 2, 4, 6}
// gfn.RangeBy(10, 0, -2) // []int{10, 8, 6, 4, 2}
func RangeBy[T Int | Uint](start, end, step T) []T {
	if step == 0 {
		panic("step must not be zero")
	}

	if start < end && step > 0 {
		res := make([]T, 0, (end-start)/step)
		for i := start; i < end; i += step {
			res = append(res, i)
		}
		return res
	}
	if start > end && step < 0 {
		res := make([]T, 0, (end-start)/step)
		for i := start; i > end; i += step {
			res = append(res, i)
		}
		return res
	}
	return []T{}
}

// Shuffle randomizes the order of elements by using Fisher–Yates algorithm
// @example
// array := []int{1, 2, 3, 4}
// gfn.Shuffle(array)
// // array: []int{2, 1, 4, 3} or other random order
func Shuffle[T any](array []T) {
	for i := range array {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

// Equal returns true if two arrays are equal. Two arrays are considered equal
// if both are nil, or if their lengths are equal and their elements are equal.
// Elements are compared using == operator.
// @example
// gfn.Equal([]int{1, 2, 3}, []int{1, 2, 3})                    // true
// gfn.Equal([]string{"a", "c", "b"}, []string{"a", "b", "c"})  // false
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aa := range a {
		if aa != b[i] {
			return false
		}
	}
	return true
}

/* @example EqualBy
a := []int{1, 2, 3, 4, 5}
b := []rune{'a', 'b', 'c', 'd', 'e'}
gfn.EqualBy(a, b, func(aa int, bb rune) bool {
	return (aa - 1) == int(bb-'a')
}) // true
*/

// EqualBy returns true if two arrays are equal by comparing their elements
// using the given function.
func EqualBy[T1, T2 any](a []T1, b []T2, fn func(T1, T2) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aa := range a {
		if !fn(aa, b[i]) {
			return false
		}
	}
	return true
}

// ToSet converts an array to a set.
// @example
// gfn.ToSet([]int{0, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5})
// // map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}}
func ToSet[T comparable](array []T) map[T]struct{} {
	res := make(map[T]struct{})
	for _, v := range array {
		res[v] = struct{}{}
	}
	return res
}

// IsSortedBy returns true if the array is sorted in the given order.
// The order function should return true if a1 is ok to be placed before a2.
// @example
// gfn.IsSortedBy([]int{2, 2, 1, 1, -1, -1}, func(a, b int) bool { return a >= b })
// // true
func IsSortedBy[T any](array []T, order func(a1, a2 T) bool) bool {
	for i := 0; i < len(array)-1; i++ {
		if !order(array[i], array[i+1]) {
			return false
		}
	}
	return true
}

// IsSorted returns true if the array is sorted in ascending order.
// @example
// gfn.IsSorted([]int{1, 2, 3, 4})  // true
func IsSorted[T Int | Uint | Float | ~string](array []T) bool {
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			return false
		}
	}
	return true
}

// Counter returns a map of values and their counts.
// @example
// gfn.Counter([]int{1, 2, 2, 2, 2})  // map[int]int{1: 1, 2: 4}
func Counter[T comparable](array []T) map[T]int {
	res := make(map[T]int)
	for _, v := range array {
		res[v]++
	}
	return res
}

/* @example CounterBy
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
*/

// CounterBy returns a map of values and their counts. The values are
// calculated by the given function.
func CounterBy[T any, U comparable](array []T, fn func(T) U) map[U]int {
	res := make(map[U]int)
	for _, v := range array {
		res[fn(v)]++
	}
	return res
}

/* @example Zip
gfn.Zip([]int{1, 2, 3}, []string{"a", "b", "c"})
// []gfn.Pair[int, string]{
// 	{First: 1, Second: "a"},
// 	{First: 2, Second: "b"},
// 	{First: 3, Second: "c"}
// }
*/

// Zip returns a sequence of pairs built from the elements of two arrays.
func Zip[T, U any](a []T, b []U) []Pair[T, U] {
	l := Min(len(a), len(b))
	res := make([]Pair[T, U], l)
	for i := 0; i < l; i++ {
		res[i] = Pair[T, U]{a[i], b[i]}
	}
	return res
}

/* @example Unzip
pairs := []gfn.Pair[int, string]{
	{First: 1, Second: "a"},
	{First: 2, Second: "b"},
	{First: 3, Second: "c"},
}
gfn.Unzip(len(pairs), func(i int) (int, string) {
	return pairs[i].First, pairs[i].Second
})
// ([]int{1, 2, 3}, []string{"a", "b", "c"})
*/

// Unzip returns two arrays built from the elements of a sequence of pairs.
func Unzip[T, U any](n int, unzipFn func(i int) (T, U)) ([]T, []U) {
	if n < 0 {
		panic("negative length")
	}
	a := make([]T, n)
	b := make([]U, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = unzipFn(i)
	}
	return a, b
}

// Sample returns a random sample of n elements from an array. Every position in
// the array are at most selected once. n should be less or equal to len(array).
// @example
// gfn.Sample([]int{1, 2, 3, 4, 5}, 3)  // []int{3, 1, 5} or other random choices.
func Sample[T any](array []T, n int) []T {
	if n < 0 {
		panic("negative length")
	}
	if n > len(array) {
		panic("sample size larger than array length")
	}
	indexes := Range(0, n)
	Shuffle(indexes)
	res := make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = array[indexes[i]]
	}
	return res
}

// Uniq returns an array with all duplicates removed.
// @example
// gfn.Uniq([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})  // []int{1, 2, 3, 4}
func Uniq[T comparable](array []T) []T {
	res := []T{}
	seen := make(map[T]struct{})
	for _, v := range array {
		if _, ok := seen[v]; !ok {
			res = append(res, v)
			seen[v] = struct{}{}
		}
	}
	return res
}

/* @example UniqBy
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
*/

// UniqBy returns an array with all duplicates removed by applying a function to each element.
func UniqBy[T any, U comparable](array []T, fn func(T) U) []T {
	res := []T{}
	seen := make(map[U]struct{})
	for _, v := range array {
		value := fn(v)
		if _, ok := seen[value]; !ok {
			res = append(res, v)
			seen[value] = struct{}{}
		}
	}
	return res
}

// Union returns an array with all duplicates removed from multiple arrays.
// @example
// gfn.Union([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
// // []int{1, 2, 3, 4, 5}
func Union[T comparable](arrays ...[]T) []T {
	res := []T{}
	seen := make(map[T]struct{})
	for _, array := range arrays {
		for _, v := range array {
			if _, ok := seen[v]; !ok {
				res = append(res, v)
				seen[v] = struct{}{}
			}
		}
	}
	return res
}

/* @example UnionBy
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
// 	{"Alice", "Accounting"},
// 	{"Bob", "Accounting"},
// 	{"Cindy", "Engineering"},
// 	{"Dave", "Engineering"},
// 	{"Eve", "Engineering"},
// }
*/

// UnionBy returns an array with all duplicates removed from multiple arrays
// by applying a function to each element.
func UnionBy[T any, U comparable](fn func(T) U, arrays ...[]T) []T {
	res := []T{}
	seen := make(map[U]struct{})
	for _, array := range arrays {
		for _, v := range array {
			value := fn(v)
			if _, ok := seen[value]; !ok {
				res = append(res, v)
				seen[value] = struct{}{}
			}
		}
	}
	return res
}

// Copy returns a new array that is a shallow copy of the original array.
// @example
// gfn.Copy([]int{1, 2, 3})  // []int{1, 2, 3}
//
// array := []int{1, 2, 3, 4, 5, 6}
// gfn.Copy(array[2:])
// // []int{3, 4, 5, 6}
func Copy[T any](array []T) []T {
	res := make([]T, len(array))
	copy(res, array)
	return res
}

// Difference returns a new array that is a copy of the original array,
// removing all occurrences of any item that also appear in others.
// The order is preserved from the original array.
// @example
// gfn.Difference([]int{1, 2, 3, 4}, []int{2, 4})  // []int{1, 3}
func Difference[T comparable](array []T, others ...[]T) []T {
	res := Copy(array)
	for _, other := range others {
		m := ToSet(other)
		res = Filter(res, func(v T) bool {
			_, ok := m[v]
			return !ok
		})
	}
	return res
}

/* @example DifferenceBy
type Data struct {
	value int
}
data1 := []Data{{1}, {3}, {2}, {4}, {5}, {2}}
data2 := []Data{{3}, {4}, {5}}
gfn.DifferenceBy(func(d Data) int { return d.value }, data1, data2)
// []Data{{1}, {2}, {2}}
*/

// DifferenceBy returns a new array that is a copy of the original array,
// removing all occurrences of any item that also appear in others. The occurrences
// are determined by applying a function to each element.
func DifferenceBy[T any, U comparable](fn func(T) U, array []T, others ...[]T) []T {
	res := make([]Pair[U, T], len(array))
	for i, v := range array {
		res[i] = Pair[U, T]{fn(v), v}
	}
	for _, other := range others {
		seen := map[U]struct{}{}
		for _, v := range other {
			seen[fn(v)] = struct{}{}
		}
		res = Filter(res, func(p Pair[U, T]) bool {
			_, ok := seen[p.First]
			return !ok
		})
	}
	return Map(res, func(p Pair[U, T]) T {
		return p.Second
	})
}

// Fill sets all elements of an array to a given value.
// You can control the start and end index by using the slice.
// @example
// array := make([]bool, 5)
// gfn.Fill(array, true)
// // []bool{true, true, true, true, true}
//
// // you can control the start and end index by using the slice
// array2 := make([]int, 5)
// gfn.Fill(array2[2:], 100)
// // []int{0, 0, 100, 100, 100}
func Fill[T any](array []T, value T) {
	for i := range array {
		array[i] = value
	}
}

// Count returns the number of occurrences of a value in an array.
// @example
// gfn.Count([]int{1, 2, 2, 2, 5, 6}, 2)  // 3
func Count[T comparable](array []T, value T) int {
	res := 0
	for _, v := range array {
		if v == value {
			res++
		}
	}
	return res
}

/* @example CountBy
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
*/

// CountBy returns the number of elements in an array that satisfy a predicate.
func CountBy[T any](array []T, fn func(T) bool) int {
	res := 0
	for _, v := range array {
		if fn(v) {
			res++
		}
	}
	return res
}

/* @example GroupBy
array := []int{1, 2, 3, 4, 5, 6, 7, 8}
groups := gfn.GroupBy(array, func(i int) string {
	if i%2 == 0 {
		return "even"
	}
	return "odd"
})
// map[string][]int{
// 	"even": []int{2, 4, 6, 8},
// 	"odd":  []int{1, 3, 5, 7},
// }
*/

// GroupBy generate a map of arrays by grouping the elements of an array
// according to a given function.
func GroupBy[T any, K comparable](array []T, groupFn func(T) K) map[K][]T {
	res := make(map[K][]T)
	for _, v := range array {
		k := groupFn(v)
		res[k] = append(res[k], v)
	}
	return res
}

// IndexOf returns the index of the first occurrence of a value in an array,
// or -1 if not found.
// @example
// gfn.IndexOf([]int{1, 2, 3, 4}, 3)  // 2
// gfn.IndexOf([]int{1, 2, 3, 4}, 5)  // -1
func IndexOf[T comparable](array []T, value T) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of a value in an array,
// or -1 if not found.
// @example
// gfn.LastIndexOf([]int{3, 3, 3, 4}, 3)  // 2
// gfn.LastIndexOf([]int{1, 2, 3, 4}, 5)  // -1
func LastIndexOf[T comparable](array []T, value T) int {
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == value {
			return i
		}
	}
	return -1
}

// Reverse reverses an array in place.
// @example
// array := []int{1, 2, 3, 4}
// gfn.Reverse(array)
// // []int{4, 3, 2, 1}
func Reverse[T any](array []T) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

/* @example All
gfn.All([]int{1, 2, 3, 4}, func(i int) bool {
	return i > 0
}
// true
*/

// All returns true if all elements in an array pass a given test.
func All[T any](array []T, fn func(T) bool) bool {
	for _, v := range array {
		if !fn(v) {
			return false
		}
	}
	return true
}

/* @example Any
gfn.Any([]int{1, 2, 3, 4}, func(i int) bool {
	return i > 3
}
// true
*/

// Any returns true if at least one element in an array passes a given test.
func Any[T any](array []T, fn func(T) bool) bool {
	for _, v := range array {
		if fn(v) {
			return true
		}
	}
	return false
}

// Concat returns a new array that is the result of joining two or more arrays.
// @example
// gfn.Concat([]int{1, 2}, []int{3, 4})  // []int{1, 2, 3, 4}
func Concat[T any](arrays ...[]T) []T {
	var res []T
	for _, array := range arrays {
		res = append(res, array...)
	}
	return res
}

/* @example Find
value, index := gfn.Find([]string{"a", "ab", "abc"}, func(s string) bool {
	return len(s) > 1
})
// "ab", 1
*/

// Find returns the first element in an array that passes a given test and corresponding index.
// Index of -1 is returned if no element passes the test.
func Find[T any](array []T, fn func(T) bool) (T, int) {
	for i, v := range array {
		if fn(v) {
			return v, i
		}
	}
	var res T
	return res, -1
}

/* @example FindLast
value, index := gfn.FindLast([]string{"a", "ab", "abc"}, func(s string) bool {
	return len(s) > 1
})
// "abc", 2
*/

// FindLast returns the last element in an array that passes a given test and corresponding index.
// Index of -1 is returned if no element passes the test.
func FindLast[T any](array []T, fn func(T) bool) (T, int) {
	for i := len(array) - 1; i >= 0; i-- {
		if fn(array[i]) {
			return array[i], i
		}
	}
	var res T
	return res, -1
}

// Remove removes all elements from an array that equal to given values.
// @example
// gfn.Remove([]int{1, 2, 3, 4, 2, 3, 2, 3}, 2, 3)  // []int{1, 4}
func Remove[T comparable](array []T, values ...T) []T {
	res := []T{}
	valueSet := ToSet(values)
	for _, v := range array {
		_, ok := valueSet[v]
		if !ok {
			res = append(res, v)
		}
	}
	return res
}

// Intersection returns a new array that is the intersection of two or more arrays.
// @example
// arr1 := []int{1, 2, 3, 4, 5}
// arr2 := []int{2, 3, 4, 5, 6}
// arr3 := []int{5, 4, 3, 2}
// arr4 := []int{2, 3}
// gfn.Intersection(arr1, arr2, arr3, arr4)  // []int{2, 3}
func Intersection[T comparable](arrays ...[]T) []T {
	if len(arrays) <= 1 {
		panic("requires at least 2 arrays")
	}

	res := Uniq(arrays[0])
	for _, arr := range arrays[1:] {
		set := ToSet(arr)
		res = Filter(res, func(v T) bool {
			_, ok := set[v]
			return ok
		})
	}
	return res
}

/* @example IntersectionBy
type Data struct {
	value int
}
data1 := []Data{{1}, {3}, {2}, {4}, {5}}
data2 := []Data{{2}, {3}}
gfn.IntersectionBy(func(d Data) int { return d.value }, data1, data2)
// []Data{{3}, {2}}
*/

// IntersectionBy returns a new array that is the intersection of two or more arrays,
// where intersection is determined by a given function.
func IntersectionBy[T any, U comparable](fn func(T) U, arrays ...[]T) []T {
	if len(arrays) <= 1 {
		panic("requires at least 2 arrays")
	}
	// make unique pair of array[0]
	var res []Pair[U, T]
	seen := map[U]struct{}{}
	for _, v := range arrays[0] {
		key := fn(v)
		if _, ok := seen[key]; !ok {
			res = append(res, Pair[U, T]{key, v})
			seen[key] = struct{}{}
		}
	}
	// filter by seen
	for _, arr := range arrays[1:] {
		seen = map[U]struct{}{}
		for _, v := range arr {
			seen[fn(v)] = struct{}{}
		}
		res = Filter(res, func(p Pair[U, T]) bool {
			_, ok := seen[p.First]
			return ok
		})
	}
	return Map(res, func(p Pair[U, T]) T {
		return p.Second
	})
}

// Repeat returns a new array that is the result of repeating an array
// a given number of times.
// @example
// gfn.Repeat([]int{1, 2, 3}, 3)  // []int{1, 2, 3, 1, 2, 3, 1, 2, 3}
func Repeat[T any](array []T, repeat int) []T {
	if repeat < 0 {
		panic("repeat must be greater or equal to 0")
	}
	if repeat == 0 {
		return []T{}
	}
	if repeat == 1 {
		return Copy(array)
	}

	res := make([]T, len(array)*repeat)
	for i := 0; i < repeat; i++ {
		copy(res[i*len(array):], array)
	}
	return res
}

/* @example ForEach
sum := 0
gfn.ForEach([]int{1, 2, 3}, func(i int) {
	sum += i
})
// sum == 6
*/

// ForEach executes a provided function once for each array element.
func ForEach[T any](array []T, fn func(value T)) {
	for _, v := range array {
		fn(v)
	}
}

// Chunk splits an array into chunks of given size.
// @example
// gfn.Chunk([]int{1, 2, 3, 4, 5}, 2)  // [][]int{{1, 2}, {3, 4}, {5}}
func Chunk[T any](array []T, size int) [][]T {
	if size <= 0 {
		panic("size must be greater than 0")
	}

	var res [][]T
	for i := 0; i < len(array); i += size {
		end := i + size
		if end > len(array) {
			end = len(array)
		}
		res = append(res, array[i:end])
	}
	return res
}
