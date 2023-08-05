package gfn

// Compare returns true if two maps/sets are equal.
// @example
// map1 := map[int]struct{}{1: {}, 2: {}, 3: {}}
// map2 := map[int]struct{}{1: {}, 2: {}, 3: {}}
// Compare(map1, map2) // true
func Compare[T, V comparable](a map[T]V, b map[T]V) bool {
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
