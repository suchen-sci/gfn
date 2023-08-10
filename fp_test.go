package gfn_test

import (
	"strconv"
	"testing"

	. "github.com/suchen-sci/gfn"
)

func TestMap(t *testing.T) {
	{
		// int to string
		actual := Map([]int{1, 2, 3}, func(i int) string {
			return strconv.Itoa(i)
		})
		AssertSliceEqual(t, []string{"1", "2", "3"}, actual)
	}
	{
		// string to struct
		type data struct {
			data string
		}
		actual := Map([]string{"a", "b", "c"}, func(s string) data {
			return data{data: s}
		})
		AssertSliceEqual(t, []data{{data: "a"}, {data: "b"}, {data: "c"}}, actual)
	}
	{
		// string to string
		actual := Map([]string{"a", "b", "c"}, func(s string) string {
			return s + s
		})
		AssertSliceEqual(t, []string{"aa", "bb", "cc"}, actual)
	}
}

func TestFilter(t *testing.T) {
	AssertSliceEqual(t, []int{2, 4, 6}, Filter([]int{1, 2, 3, 4, 5, 6}, func(i int) bool {
		return i%2 == 0
	}))

	AssertSliceEqual(t, []string{"abc", "def"}, Filter([]string{"abc", "def", "abcdef"}, func(s string) bool {
		return len(s) <= 3
	}))

	{
		type data struct {
			data int
		}
		expected := []data{{data: 2}, {data: 4}, {data: 6}}
		inputData := []data{{data: 1}, {data: 2}, {data: 3}, {data: 4}, {data: 5}, {data: 6}}
		AssertSliceEqual(t, expected, Filter(inputData, func(d data) bool {
			return d.data%2 == 0
		}))
	}
}

func TestReduce(t *testing.T) {
	AssertEqual(t, 6, Reduce([]int{1, 2, 3}, 0, func(a, b int) int {
		return a + b
	}))
}

func TestFilterKV(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	m = FilterKV(m, func(k int, v string) bool {
		return k == 1 || v == "c"
	})
	AssertMapEqual(t, map[int]string{1: "a", 3: "c"}, m)
}

func TestReduceKV(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	total := ReduceKV(m, 0, func(value int, k string, v int) int {
		return value + v
	})
	AssertEqual(t, 6, total)
}
