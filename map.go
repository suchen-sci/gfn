package gfn

// Same returns true if two maps/sets are equal.
func Same[T, V comparable](a map[T]V, b map[T]V) bool {
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
