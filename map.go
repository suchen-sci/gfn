package gfn

/* @example EqualKV
map1 := map[int]struct{}{1: {}, 2: {}, 3: {}}
map2 := map[int]struct{}{1: {}, 2: {}, 3: {}}
gfn.EqualKV(map1, map2) // true
*/

// EqualKV returns true if two maps/sets are equal.
func EqualKV[K, V comparable](a map[K]V, b map[K]V) bool {
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

/* @example EqualKVBy
m1 := map[int]string{1: "a", 2: "b", 3: "c"}
m2 := map[int]string{1: "e", 2: "f", 3: "g"}
gfn.EqualKVBy(m1, m2, func(k int, a, b string) bool {
	return len(a) == len(b)
}) // true
*/

// EqualKVBy returns true if two maps/sets are equal by a custom function.
func EqualKVBy[K comparable, V1, V2 any](a map[K]V1, b map[K]V2, fn func(K, V1, V2) bool) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v2, ok := b[k]; !ok || !fn(k, v, v2) {
			return false
		}
	}
	return true
}

/* @example Keys
gfn.Keys(map[int]string{1: "a", 2: "b", 3: "c"})
// []int{1, 2, 3} or []int{3, 2, 1} or []int{2, 1, 3} etc.
*/

// Keys returns the keys of a map.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

/* @example Values
gfn.Values(map[int]string{1: "a", 2: "b", 3: "c"})
// []string{"a", "b", "c"} or []string{"c", "b", "a"} or []string{"b", "a", "c"} etc.
*/

// Values returns the values of a map.
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

/* @example Clear
m := map[int]string{1: "a", 2: "b", 3: "c"}
gfn.Clear(m)
// m is now an empty map
*/

// Clear removes all keys from a map.
func Clear[K comparable, V any](m map[K]V) {
	for k := range m {
		delete(m, k)
	}
}

/* @example Items
m := map[int]string{1: "a", 2: "b", 3: "c"}

gfn.Items(m)
// []gfn.Pair[int, string]{
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

/* @example Update
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
*/

// Update updates a map with the keys and values from other maps.
func Update[K comparable, V any](m map[K]V, other ...map[K]V) {
	for _, o := range other {
		for k, v := range o {
			m[k] = v
		}
	}
}

/* @example Clone
m := map[int]string{1: "a", 2: "b", 3: "c"}
m2 := gfn.Clone(m)
// m2 is a copy of m
*/

// Clone returns a shallow copy of a map.
func Clone[K comparable, V any](m map[K]V) map[K]V {
	res := make(map[K]V, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}

/* @example DeleteBy
m := map[int]string{1: "a", 2: "b", 3: "c"}
gfn.DeleteBy(m, func(k int, v string) bool {
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

/* @example Select
m := map[int]string{1: "a", 2: "b", 3: "c"}
gfn.Select(m, func(k int, v string) bool {
	return k == 1 || v == "c"
})
// map[int]string{1: "a", 3: "c"}
*/

// Select returns a map with keys and values that satisfy the predicate function.
func Select[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	res := make(map[K]V)
	for k, v := range m {
		if fn(k, v) {
			res[k] = v
		}
	}
	return res
}

/* @example IsDisjoint
m1 := map[int]string{1: "a", 2: "b", 3: "c"}
m2 := map[int]int{4: 4, 5: 5, 6: 6}
IsDisjoint(m1, m2)  // true

m3 := map[int]struct{}{1: {}, 2: {}, 3: {}}
m4 := map[int]struct{}{4: {}, 5: {}, 6: {}}
gfn.IsDisjoint(m3, m4)  // true
*/

// IsDisjoint returns true if the maps have no keys in common. It usually
// used to check if two sets are disjoint.
func IsDisjoint[K comparable, V1 any, V2 any](m1 map[K]V1, m2 map[K]V2) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1 := range m1 {
		if _, ok := m2[k1]; ok {
			return false
		}
	}
	return true
}

/* @example IntersectKeys
m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
m2 := map[int]string{1: "a", 2: "b"}
m3 := map[int]string{2: "b", 3: "c", 4: "d"}
gfn.IntersectKeys(m1, m2, m3)  // []int{2}
*/

// IntersectKeys returns a slice of keys that are in all maps. It usually
// used to find the intersection of two or more sets.
func IntersectKeys[K comparable, V any](ms ...map[K]V) []K {
	if len(ms) == 0 {
		return nil
	}
	if len(ms) == 1 {
		return Keys(ms[0])
	}

	res := Keys(ms[0])
	for _, m := range ms[1:] {
		res = Filter(res, func(k K) bool {
			_, ok := m[k]
			return ok
		})
	}
	return res
}

/* @example DifferentKeys
m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
m2 := map[int]string{1: "a", 2: "b"}
m3 := map[int]string{2: "b", 3: "c"}
gfn.DifferentKeys(m1, m2, m3)  // []int{4}
*/

// DifferentKeys returns a slice of keys that are in the first map but not in the others,
// only keys in the map are considered, not values. It usually used to find the
// difference between two or more sets.
func DifferentKeys[K comparable, V any](ms ...map[K]V) []K {
	if len(ms) == 0 {
		return nil
	}
	if len(ms) == 1 {
		return Keys(ms[0])
	}

	res := Keys(ms[0])
	for _, m := range ms[1:] {
		res = Filter(res, func(k K) bool {
			_, ok := m[k]
			return !ok
		})
	}
	return res
}

/* @example GetOrDefault
m := map[int]string{1: "a", 2: "b", 3: "c"}
gfn.GetOrDefault(m, 1, "d")  // "a"
gfn.GetOrDefault(m, 4, "d")  // "d"
*/

// GetOrDefault returns the value for a key if it exists, otherwise it returns the default value.
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	value, ok := m[key]
	if ok {
		return value
	}
	return defaultValue
}

/* @example ForEachKV
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
*/

// ForEachKV iterates over a map and calls a function for each key/value pair.
func ForEachKV[K comparable, V any](m map[K]V, fn func(K, V)) {
	for k, v := range m {
		fn(k, v)
	}
}

/* @example ToKV
gfn.ToKV(3, func(i int) (int, string) {
	return i, strconv.Itoa(i)
})
// map[int]string{0: "0", 1: "1", 2: "2"}
*/

// ToKV converts a slice to a map using a function to generate the keys and values.
func ToKV[K comparable, V any](n int, fn func(int) (K, V)) map[K]V {
	if n < 0 {
		panic("n must be greater than or equal to zero")
	}
	m := make(map[K]V, n)
	for i := 0; i < n; i++ {
		k, v := fn(i)
		m[k] = v
	}
	return m
}
