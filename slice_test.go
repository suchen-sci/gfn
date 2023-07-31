package gfn_test

import (
	"testing"

	. "github.com/suchen-sci/gfn"
)

func TestContains(t *testing.T) {
	AssertTrue(t, Contains([]int{1, 2, 3}, 2))
	AssertFalse(t, Contains([]int{1, 2, 3}, 4))

	AssertTrue(t, Contains([]string{"a", "b", "c"}, "b"))
	AssertFalse(t, Contains([]string{"a", "b", "c"}, "d"))

	AssertTrue(t, Contains([]bool{true, false, true}, true))
	AssertFalse(t, Contains([]bool{true, true, true}, false))

	AssertTrue(t, Contains([]float64{1.1, 2.2, 3.3}, 2.2))
	AssertFalse(t, Contains([]float64{1.1, 2.2, 3.3}, 4.4))

	type data struct {
		data int
	}
	AssertTrue(t, Contains([]data{{data: 1}, {data: 2}}, data{data: 2}))
	AssertFalse(t, Contains([]data{{data: 1}, {data: 2}}, data{data: 3}))

	dataA := &data{data: 1}
	dataB := &data{data: 1}
	AssertTrue(t, Contains([]*data{dataA, dataB}, dataA))
	AssertFalse(t, Contains([]*data{dataA, dataB}, &data{data: 1}))
}
