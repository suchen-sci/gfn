package gfn_test

import (
	"sort"
	"strconv"
	"testing"
	"time"

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

	AssertTrue(t, Contains([]time.Duration{time.Second, 2 * time.Second}, time.Second))
	AssertFalse(t, Contains([]time.Duration{time.Second, 2 * time.Second}, 3*time.Second))
}

func TestRange(t *testing.T) {
	AssertSliceEqual(t, []int{0, 1, 2, 3, 4, 5, 6}, Range(0, 7))
	AssertSliceEqual(t, []int{5, 6, 7, 8, 9}, Range(5, 10))
	AssertSliceEqual(t, []int{}, Range(0, 0))
	AssertSliceEqual(t, []int{}, Range(10, 0))
	AssertSliceEqual(t, []int{-3, -2, -1, 0, 1}, Range(-3, 2))

	AssertSliceEqual(t, []uint{0, 1, 2, 3, 4, 5, 6}, Range[uint](0, 7))
	AssertSliceEqual(t, []uint{5, 6, 7, 8, 9}, Range[uint](5, 10))
	AssertSliceEqual(t, []uint{}, Range[uint](0, 0))
	AssertSliceEqual(t, []uint{}, Range[uint](10, 0))
}

func TestRangeBy(t *testing.T) {
	AssertSliceEqual(t, []int{0, 1, 2, 3, 4, 5, 6}, RangeBy(0, 7, 1))
	AssertSliceEqual(t, []int{0, 3, 6, 9, 12, 15}, RangeBy(0, 17, 3))
	AssertSliceEqual(t, []int{}, RangeBy(0, 5, -2))
	AssertSliceEqual(t, []int{-10, -8, -6, -4, -2, 0}, RangeBy(-10, 1, 2))
	AssertSliceEqual(t, []int{}, RangeBy(0, 0, 1))
	AssertSliceEqual(t, []int{0, 2, 4, 6}, RangeBy(0, 8, 2))
	AssertSliceEqual(t, []int{10, 8, 6, 4, 2}, RangeBy(10, 0, -2))

	AssertSliceEqual(t, []int{13, 11, 9, 7}, RangeBy(13, 5, -2))
	AssertSliceEqual(t, []int{29, 24, 19, 14, 9, 4}, RangeBy(29, 0, -5))
	AssertSliceEqual(t, []int{}, RangeBy(5, 0, 2))
	AssertSliceEqual(t, []int{1, -1, -3, -5, -7, -9}, RangeBy(1, -10, -2))
	AssertSliceEqual(t, []int{}, RangeBy(0, 0, -1))

	AssertSliceEqual(t, []uint{0, 1, 2, 3, 4, 5, 6}, RangeBy[uint](0, 7, 1))
	AssertSliceEqual(t, []uint{0, 3, 6, 9, 12, 15}, RangeBy[uint](0, 17, 3))
	AssertSliceEqual(t, []uint{}, RangeBy[uint](0, 0, 1))
	AssertSliceEqual(t, []uint{}, RangeBy[uint](5, 0, 2))

	AssertSliceEqual(t, []int64{}, RangeBy[int64](0, 5, -2))
	AssertSliceEqual(t, []int64{-10, -8, -6, -4, -2, 0}, RangeBy[int64](-10, 1, 2))
	AssertSliceEqual(t, []int64{13, 11, 9, 7}, RangeBy[int64](13, 5, -2))
	AssertSliceEqual(t, []int64{29, 24, 19, 14, 9, 4}, RangeBy[int64](29, 0, -5))
	AssertSliceEqual(t, []int64{1, -1, -3, -5, -7, -9}, RangeBy[int64](1, -10, -2))
	AssertSliceEqual(t, []int64{}, RangeBy[int64](0, 0, -1))

	AssertPanics(t, func() {
		RangeBy(0, 10, 0)
	})
}

func TestShuffle(t *testing.T) {
	array := Range(0, 200000)
	Shuffle(array)
	AssertFalse(t, IsSorted(array))
	AssertFalse(t, Equal(Range(0, 200000), array))

	sort.Ints(array)
	AssertTrue(t, IsSorted(array))
	AssertSliceEqual(t, Range(0, 200000), array)
}

func TestEqual(t *testing.T) {
	AssertTrue(t, Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	AssertFalse(t, Equal([]int{1, 3, 2}, []int{1, 2, 3}))

	AssertTrue(t, Equal([]string{"a", "b", "c"}, []string{"a", "b", "c"}))
	AssertFalse(t, Equal([]string{"a", "c", "b"}, []string{"a", "b", "c"}))
}

func TestToSet(t *testing.T) {
	set := ToSet([]int{0, 1, 2, 3, 4, 5})
	AssertEqual(t, 6, len(set))
	for i := 0; i < 6; i++ {
		_, ok := set[i]
		AssertTrue(t, ok, strconv.Itoa(i))
	}

	set = ToSet([]int{0, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5})
	AssertEqual(t, 6, len(set))
	for i := 0; i < 6; i++ {
		_, ok := set[i]
		AssertTrue(t, ok, strconv.Itoa(i))
	}

	expected := map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}}
	AssertTrue(t, Same(expected, set))
}
