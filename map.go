package gfn

// Compare returns true if two maps/sets are equal.
// @example
// map1 := map[int]struct{}{1: {}, 2: {}, 3: {}}
// map2 := map[int]struct{}{1: {}, 2: {}, 3: {}}
// Compare(map1, map2) // true
func Compare[K, V comparable](a map[K]V, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v2, ok := b[k]; !ok || v != v2 {
			return false
		}
	}
	return true
}

// Keys returns the keys of a map.
// @example
// gfn.Keys(map[int]string{1: "a", 2: "b", 3: "c"})
// // []int{1, 2, 3} or []int{3, 2, 1} or []int{2, 1, 3} etc.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// Values returns the values of a map.
// @example
// gfn.Values(map[int]string{1: "a", 2: "b", 3: "c"})
// // []string{"a", "b", "c"} or []string{"c", "b", "a"} or []string{"b", "a", "c"} etc.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

/* @example Invert
m := map[string]string{
	"Array": "array.go",
	"Map":   "map.go",
	"Set":   "set.go",
	"Math":  "math.go",
}

gfn.Invert(m)
// map[string]string{
// 	"array.go": "Array",
// 	"map.go":   "Map",
// 	"set.go":   "Set",
// 	"math.go":  "Math",
// }
*/

// Invert returns a map with keys and values swapped.
func Invert[K, V comparable](m map[K]V) map[V]K {
	res := make(map[V]K)
	for k, v := range m {
		res[v] = k
	}
	return res
}

// Clear removes all keys from a map.
// @example
// m := map[int]string{1: "a", 2: "b", 3: "c"}
// gfn.Clear(m)
// // m is now an empty map
func Clear[K comparable, V any](m map[K]V) {
	for k := range m {
		delete(m, k)
	}
}

/* @example Items
m := map[int]string{1: "a", 2: "b", 3: "c"}

gfn.Items(m)
// []Pair[int, string]{
// 	{1, "a"},
// 	{2, "b"},
// 	{3, "c"},
// }
*/

// Items returns a slice of pairs of keys and values.
func Items[K comparable, V any](m map[K]V) []Pair[K, V] {
	items := make([]Pair[K, V], len(m))
	i := 0
	for k, v := range m {
		items[i] = Pair[K, V]{k, v}
		i++
	}
	return items
}

// Update updates a map with the keys and values from other maps.
// @example
// m1 := map[int]string{1: "a", 2: "b", 3: "c"}
// m2 := map[int]string{1: "d", 2: "e"}
// m3 := map[int]string{1: "f"}
// Update(m1, m2, m3)
// // map[int]string{1: "f", 2: "e", 3: "c"}
func Update[K comparable, V any](m map[K]V, other ...map[K]V) {
	for _, o := range other {
		for k, v := range o {
			m[k] = v
		}
	}
}

// Clone returns a shallow copy of a map.
// @example
// m := map[int]string{1: "a", 2: "b", 3: "c"}
// m2 := Clone(m)
// // m2 is a copy of m
func Clone[K comparable, V any](m map[K]V) map[K]V {
	res := make(map[K]V, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}

/* @example DeleteBy
m := map[int]string{1: "a", 2: "b", 3: "c"}
DeleteBy(m, func(k int, v string) bool {
	return k == 1 || v == "c"
})
// map[int]string{2: "b"}
*/

// DeleteBy deletes keys from a map if the predicate function returns true.
func DeleteBy[K comparable, V any](m map[K]V, deleteFn func(K, V) bool) {
	for k, v := range m {
		if deleteFn(k, v) {
			delete(m, k)
		}
	}
}

/* @example Rejected
m := map[int]string{1: "a", 2: "b", 3: "c"}
rejected := Rejected(m, func(k int, v string) bool {
	return k == 1 || v == "c"
})
// map[int]string{2: "b"}
*/

// Rejected returns a map with keys and values that don't pass the predicate function.
func Rejected[K comparable, V any](m map[K]V, keep func(K, V) bool) map[K]V {
	res := make(map[K]V)
	for k, v := range m {
		if !keep(k, v) {
			res[k] = v
		}
	}
	return res
}
