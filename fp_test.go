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
