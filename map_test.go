package gfn_test

import (
	"strconv"
	"testing"

	. "github.com/suchen-sci/gfn"
)

func TestSame(t *testing.T) {
	{
		map1 := map[int]struct{}{}
		map2 := map[int]struct{}{}
		for i := 0; i < 100; i++ {
			map1[i] = struct{}{}
			map2[i] = struct{}{}
		}
		AssertTrue(t, Same(map1, map2))

		map2[100] = struct{}{}
		AssertFalse(t, Same(map1, map2))
	}

	{
		map1 := map[int]string{}
		map2 := map[int]string{}
		for i := 0; i < 100; i++ {
			map1[i] = strconv.Itoa(i)
			map2[i] = strconv.Itoa(i)
		}
		AssertTrue(t, Same(map1, map2))

		map2[0] = "999"
		AssertFalse(t, Same(map1, map2))
	}
}
