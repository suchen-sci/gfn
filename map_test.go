package gfn_test

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	. "github.com/suchen-sci/gfn"
)

func TestEqualKV(t *testing.T) {
	{
		type Number int
		type Numbers map[Number]struct{}
		map1 := Numbers{}
		map2 := Numbers{}
		for i := 0; i < 100; i++ {
			map1[Number(i)] = struct{}{}
			map2[Number(i)] = struct{}{}
		}
		AssertMapEqual(t, map1, map2)

		map2[100] = struct{}{}
		AssertFalse(t, EqualKV(map1, map2))
	}

	{
		map1 := map[int]string{}
		map2 := map[int]string{}
		for i := 0; i < 100; i++ {
			map1[i] = strconv.Itoa(i)
			map2[i] = strconv.Itoa(i)
		}
		AssertMapEqual(t, map1, map2)

		map2[0] = "999"
		AssertFalse(t, EqualKV(map1, map2))
	}
}

func TestKeys(t *testing.T) {
	{
		keys := Keys(map[int]string{1: "a", 2: "b", 3: "c"})
		sort.Ints(keys)
		AssertSliceEqual(t, []int{1, 2, 3}, keys)
	}

	{
		type Number int
		expected := []Number{1, 2, 3}
		keys := Keys(map[Number]string{1: "a", 2: "b", 3: "c"})
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
		AssertSliceEqual(t, expected, keys)
	}
}

func TestValues(t *testing.T) {
	values := Values(map[int]string{1: "a", 2: "b", 3: "c"})
	sort.Strings(values)
	AssertSliceEqual(t, []string{"a", "b", "c"}, values)
}

func TestInvert(t *testing.T) {
	{
		m := map[string]string{
			"Array": "array.go",
			"Map":   "map.go",
			"Set":   "set.go",
			"Math":  "math.go",
		}
		expected := map[string]string{
			"array.go": "Array",
			"map.go":   "Map",
			"set.go":   "Set",
			"math.go":  "Math",
		}
		AssertMapEqual(t, expected, Invert(m))
	}

	{
		m := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}
		expected := map[int]string{
			1: "a",
			2: "b",
			3: "c",
		}
		AssertMapEqual(t, expected, Invert(m))
	}
}

func TestClear(t *testing.T) {
	m := map[int]float32{}
	for i := 0; i < 1000; i++ {
		m[i] = float32(i)
	}
	Clear(m)
	AssertEqual(t, 0, len(m))
}

func TestItems(t *testing.T) {
	items := Items(map[int]string{1: "a", 2: "b", 3: "c"})
	sort.Slice(items, func(i, j int) bool {
		return items[i].First < items[j].First
	})
	expected := []Pair[int, string]{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}
	AssertSliceEqual(t, expected, items)
}

func TestUpdate(t *testing.T) {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	m2 := map[int]string{1: "d", 2: "e"}
	m3 := map[int]string{1: "f"}
	expected := map[int]string{1: "f", 2: "e", 3: "c"}
	Update(m1, m2, m3)
	AssertMapEqual(t, expected, m1)
}

func TestClone(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	clone := Clone(m)
	AssertMapEqual(t, m, clone)
	m[1] = "d"
	AssertFalse(t, EqualKV(m, clone))
}

func TestDeleteBy(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	DeleteBy(m, func(k int, v string) bool {
		return k == 1 || v == "c"
	})
	AssertMapEqual(t, map[int]string{2: "b"}, m)
}

func TestSelect(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	rejected := Select(m, func(k int, v string) bool {
		return k == 1 || v == "c"
	})
	AssertMapEqual(t, map[int]string{1: "a", 3: "c"}, rejected)
}

func TestIsDisjoint(t *testing.T) {
	{
		m1 := map[int]string{1: "a", 2: "b", 3: "c"}
		m2 := map[int]int{4: 4, 5: 5, 6: 6}
		AssertTrue(t, IsDisjoint(m1, m2))
	}
	{
		m1 := map[int]string{1: "a", 2: "b", 3: "c"}
		m2 := map[int]string{1: "a", 2: "b"}
		AssertFalse(t, IsDisjoint(m1, m2))
	}
	{
		m1 := map[int]struct{}{1: {}, 2: {}, 3: {}}
		m2 := map[int]struct{}{4: {}, 5: {}, 6: {}}
		AssertTrue(t, IsDisjoint(m1, m2))
	}
}

func TestIntersectKeys(t *testing.T) {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	m2 := map[int]string{1: "a", 2: "b"}
	m3 := map[int]string{2: "b", 3: "c", 4: "d"}
	keys := IntersectKeys(m1, m2, m3)
	AssertSliceEqual(t, []int{2}, keys)
}

func TestDifferentKeys(t *testing.T) {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	m2 := map[int]string{1: "a", 2: "b"}
	m3 := map[int]string{2: "b", 3: "c"}
	keys := DifferentKeys(m1, m2, m3)
	AssertSliceEqual(t, []int{4}, keys)
}

func TestGetOrDefault(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	AssertEqual(t, "a", GetOrDefault(m, 1, "d"))
	AssertEqual(t, "d", GetOrDefault(m, 4, "d"))
}

func TestEqualKVBy(t *testing.T) {
	{
		m1 := map[int]string{1: "a", 2: "b", 3: "c"}
		m2 := map[int]string{1: "e", 2: "f", 3: "g"}
		AssertTrue(t, EqualKVBy(m1, m2, func(k int, a, b string) bool {
			return len(a) == len(b)
		}))
	}
	{
		m1 := map[int]string{1: "a", 2: "b", 3: "c"}
		m2 := map[int]string{1: "e", 2: "f", 3: "g"}
		AssertFalse(t, EqualKVBy(m1, m2, func(k int, a, b string) bool {
			return a == b
		}))
	}

}

func TestForEachKV(t *testing.T) {
	arr := []string{}
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	ForEachKV(m, func(k int, v string) {
		arr = append(arr, fmt.Sprintf("%d:%s", k, v))
	})
	sort.Strings(arr)
	AssertSliceEqual(t, []string{"1:a", "2:b", "3:c"}, arr)
}

func TestToKV(t *testing.T) {
	m := ToKV(3, func(i int) (int, string) {
		return i, strconv.Itoa(i)
	})
	AssertMapEqual(t, map[int]string{0: "0", 1: "1", 2: "2"}, m)
}
