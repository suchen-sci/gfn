package gfn_test

import (
	"math/rand"
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

	type Number int
	type Numbers []Number
	AssertTrue(t, Contains(Numbers{1, 2, 3}, Number(2)))
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
	AssertFalse(t, Equal([]string{"a", "c", "b"}, []string{"a", "b"}))
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
	AssertMapEqual(t, expected, set)
}

func TestIsSortedBy(t *testing.T) {
	AssertTrue(t, IsSortedBy([]int{}, func(a, b int) bool { return a < b }))
	AssertTrue(t, IsSortedBy([]int{1, 2, 3, 4, 5, 6, 7}, func(a, b int) bool { return a <= b }))
	AssertTrue(t, IsSortedBy([]int{1, 1, 1, 2, 2, 2, 2}, func(a, b int) bool { return a <= b }))
	AssertTrue(t, IsSortedBy([]int{2, 2, 2, 1, 1, 1, -1, -1}, func(a, b int) bool { return a >= b }))

	AssertFalse(t, IsSortedBy([]int{1, 2, 10, 4, 5, 6, 7}, func(a, b int) bool { return a <= b }))
	AssertFalse(t, IsSortedBy([]int{1, 1, 1, 100, 2, 2, 2}, func(a, b int) bool { return a <= b }))
	AssertFalse(t, IsSortedBy([]int{2, 2, -10, 1, 1, 1, -1, -1}, func(a, b int) bool { return a >= b }))
}

func TestDistribution(t *testing.T) {
	// check empty array
	AssertMapEqual(t, map[int]int{}, Counter([]int{}))

	// check array with many elements
	{
		array := make([]int, 100000)
		for i := 0; i < 100000; i++ {
			array[i] = i
		}
		rand.Shuffle(len(array), func(i, j int) {
			array[i], array[j] = array[j], array[i]
		})
		distr := Counter(array)
		for i := 0; i < 100000; i++ {
			AssertEqual(t, 1, distr[i])
		}
	}

	// check distribution
	AssertMapEqual(t, map[int]int{1: 1, 2: 1, 3: 1, 4: 1}, Counter([]int{1, 2, 3, 4}))
	AssertMapEqual(t, map[int]int{1: 1, 2: 2, 3: 1, 4: 1}, Counter([]int{1, 2, 3, 4, 2}))
	AssertMapEqual(t, map[int]int{1: 1, 2: 4}, Counter([]int{1, 2, 2, 2, 2}))
}

func TestIsSorted(t *testing.T) {
	AssertTrue(t, IsSorted([]int{}))
	AssertTrue(t, IsSorted([]int{1, 2, 3}))
	AssertTrue(t, IsSorted([]int{1, 1, 1, 1, 1, 1}))
	AssertTrue(t, IsSorted([]int{1, 2, 2, 3, 3, 3}))

	AssertFalse(t, IsSorted([]int{1, 23, 2}))
	AssertFalse(t, IsSorted([]int{1, 23, 99, 1, 100, 2}))
}

func TestZip(t *testing.T) {
	AssertSliceEqual(t, []Pair[int, int]{}, Zip([]int{}, []int{}))
	AssertSliceEqual(t, []Pair[int, string]{{1, "a"}, {2, "b"}, {3, "c"}}, Zip([]int{1, 2, 3}, []string{"a", "b", "c"}))
	AssertSliceEqual(t, []Pair[int, string]{{1, "a"}, {2, "b"}}, Zip([]int{1, 2}, []string{"a", "b", "c"}))
}

func TestUnzip(t *testing.T) {
	pairs := []Pair[int, string]{
		{First: 1, Second: "a"},
		{First: 2, Second: "b"},
		{First: 3, Second: "c"},
	}
	a, b := Unzip(len(pairs), func(i int) (int, string) {
		return pairs[i].First, pairs[i].Second
	})
	AssertSliceEqual(t, []int{1, 2, 3}, a)
	AssertSliceEqual(t, []string{"a", "b", "c"}, b)

	AssertPanics(t, func() {
		Unzip(-1, func(i int) (int, string) {
			return 0, ""
		})
	})
}

func TestSample(t *testing.T) {
	for i := 0; i < 100; i++ {
		array := Range(i, i+100)
		samples := Sample(array, 90)
		AssertEqual(t, 90, len(samples))
		m := map[int]struct{}{}
		for _, sample := range samples {
			_, ok := m[sample]
			AssertFalse(t, ok)
			m[sample] = struct{}{}
			AssertTrue(t, Contains(array, sample))
		}
	}

	AssertPanics(t, func() {
		Sample([]int{}, 1)
	})

	AssertPanics(t, func() {
		Sample([]int{1}, -1)
	})
}

func TestUniq(t *testing.T) {
	for i := 0; i < 100; i++ {
		array := []int{}
		start := rand.Intn(100)
		for j := start; j < start+100; j++ {
			array = append(array, j, j, j)
		}
		Shuffle(array)
		array = Uniq(array)
		sort.Ints(array)
		AssertEqual(t, 100, len(array))
		AssertSliceEqual(t, Range(start, start+100), array)
	}
}

func TestUnion(t *testing.T) {
	a1 := []int{1, 2, 2, 3, 3, 4, 5}
	a2 := []int{2, 3, 3, 4, 4, 5, 6}
	a3 := []int{3, 4, 4, 5, 5, 6, 7}
	AssertSliceEqual(t, []int{1, 2, 3, 4, 5, 6, 7}, Union(a1, a2, a3))
}

func TestCopy(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := Copy(a)
	AssertSliceEqual(t, a, b)
	a[0] = 100
	AssertEqual(t, 1, b[0])

	c := []int{1, 2, 3, 4, 5}
	d := Copy(c[2:])
	AssertSliceEqual(t, []int{3, 4, 5}, d)
}

func TestDiff(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{2, 4, 6}
	c := []int{1, 3, 5, 7}
	AssertSliceEqual(t, c, Difference(a, b))

	for i := 0; i < 100; i++ {
		a := Range(i, i+200)
		b := Range(i+50, i+150)
		c := Range(i+150, i+200)
		Shuffle(a)
		Shuffle(b)
		Shuffle(c)

		d := Difference(a, b, c)
		sort.Ints(d)
		AssertSliceEqual(t, Range(i, i+50), d)
	}
}

func TestFill(t *testing.T) {
	array := make([]bool, 5)
	AssertSliceEqual(t, []bool{false, false, false, false, false}, array)
	Fill(array, true)
	AssertSliceEqual(t, []bool{true, true, true, true, true}, array)

	array2 := make([]int, 5)
	Fill(array2[2:], 100)
	AssertSliceEqual(t, []int{0, 0, 100, 100, 100}, array2)
}

func TestCount(t *testing.T) {
	array := Range(0, 200)
	Shuffle(array)
	for i := 0; i < 200; i++ {
		AssertEqual(t, 1, Count(array, i))
	}

	array2 := []int{1, 1, 1, 2, 1, 4, 1}
	AssertEqual(t, 5, Count(array2, 1))
}

func TestGroupBy(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	groups := GroupBy(array, func(i int) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})
	AssertEqual(t, 2, len(groups))
	AssertSliceEqual(t, []int{2, 4, 6, 8}, groups["even"])
	AssertSliceEqual(t, []int{1, 3, 5, 7}, groups["odd"])
}

func TestIndexOf(t *testing.T) {
	AssertEqual(t, -1, IndexOf([]int{}, 1))
	AssertEqual(t, -1, IndexOf([]int{1, 2, 3}, 4))
	AssertEqual(t, 2, IndexOf([]int{1, 2, 3}, 3))
	AssertEqual(t, 2, IndexOf([]int{1, 2, 3, 3, 3}, 3))
}

func TestLastIndexOf(t *testing.T) {
	AssertEqual(t, -1, LastIndexOf([]int{}, 1))
	AssertEqual(t, -1, LastIndexOf([]int{1, 2, 3}, 4))
	AssertEqual(t, 2, LastIndexOf([]int{3, 3, 3}, 3))
	AssertEqual(t, 0, LastIndexOf([]int{3, 2, 2, 2, 2}, 3))
}

func TestReverse(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	Reverse(array)
	AssertSliceEqual(t, []int{5, 4, 3, 2, 1}, array)

	array2 := []int{1, 2, 3, 4}
	Reverse(array2)
	AssertSliceEqual(t, []int{4, 3, 2, 1}, array2)
}

func TestAll(t *testing.T) {
	AssertTrue(t, All([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i > 0
	}))
	AssertFalse(t, All([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i > 1
	}))
}

func TestAny(t *testing.T) {
	AssertTrue(t, Any([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i > 4
	}))
	AssertFalse(t, Any([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i > 5
	}))
}

func TestConcat(t *testing.T) {
	AssertSliceEqual(t, []int{1, 2, 3, 4, 5}, Concat([]int{1, 2}, []int{3, 4, 5}))
	AssertSliceEqual(t, []int{1, 2, 3, 4, 5}, Concat([]int{1, 2}, []int{3}, []int{4, 5}))
	AssertSliceEqual(t, []int{1, 2, 3, 4, 5}, Concat([]int{1, 2}, []int{3}, []int{4}, []int{5}))
	AssertSliceEqual(t, []int{1, 2, 3, 4, 5}, Concat([]int{1, 2}, []int{}, []int{3}, []int{4}, []int{}, []int{5}))
}

func TestFind(t *testing.T) {
	{
		value, index := Find([]string{}, func(s string) bool {
			return len(s) != 0
		})
		AssertEqual(t, -1, index)
		AssertEqual(t, "", value)
	}
	{
		value, index := Find([]string{"a", "ab", "abc"}, func(s string) bool {
			return len(s) > 1
		})
		AssertEqual(t, 1, index)
		AssertEqual(t, "ab", value)
	}
}

func TestFindLast(t *testing.T) {
	{
		value, index := FindLast([]string{}, func(s string) bool {
			return len(s) != 0
		})
		AssertEqual(t, -1, index)
		AssertEqual(t, "", value)
	}
	{
		value, index := FindLast([]string{"a", "ab", "abc"}, func(s string) bool {
			return len(s) > 1
		})
		AssertEqual(t, 2, index)
		AssertEqual(t, "abc", value)
	}
}

func TestRemove(t *testing.T) {
	{
		array := []int{1, 2, 3, 4, 2, 4}
		AssertSliceEqual(t, []int{1, 3, 4, 4}, Remove(array, 2))
	}
	{
		array := []int{1, 2, 3, 4, 2, 4}
		AssertSliceEqual(t, []int{1, 3}, Remove(array, 2, 4))
	}
}

func TestIntersection(t *testing.T) {
	{
		arr1 := []int{1, 2, 3, 4, 5}
		arr2 := []int{2, 3, 4, 5, 6}
		arr3 := []int{5, 4, 3, 2}
		arr4 := []int{2, 3}
		AssertSliceEqual(t, []int{2, 3}, Intersection(arr1, arr2, arr3, arr4))
	}
	{
		arr1 := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
		arr2 := []int{2, 3, 4, 5, 6, 2, 3, 2, 3}
		arr3 := []int{5, 4, 3, 2, 2, 3}
		arr4 := []int{2, 3, 2, 3, 2, 3}
		AssertSliceEqual(t, []int{2, 3}, Intersection(arr1, arr2, arr3, arr4))
	}

	AssertPanics(t, func() {
		Intersection[int]()
	})
	AssertPanics(t, func() {
		Intersection([]int{1, 2, 3})
	})
}

func TestRepeat(t *testing.T) {
	AssertSliceEqual(t, []int{5, 5, 5}, Repeat([]int{5}, 3))
	AssertSliceEqual(t, []int{5, 3, 1, 5, 3, 1, 5, 3, 1}, Repeat([]int{5, 3, 1}, 3))
	AssertSliceEqual(t, []int{}, Repeat([]int{5, 3, 1}, 0))
	AssertSliceEqual(t, []int{5, 3, 1}, Repeat([]int{5, 3, 1}, 1))
	AssertSliceEqual(t, []int{}, Repeat[int](nil, 1))

	AssertPanics(t, func() {
		Repeat[int](nil, -1)
	})
}

func TestForEach(t *testing.T) {
	sum := 0
	ForEach([]int{1, 2, 3}, func(i int) {
		sum += i
	})
	AssertEqual(t, 6, sum)
}

func TestCountBy(t *testing.T) {
	type Employee struct {
		name       string
		department string
	}
	employees := []Employee{
		{"Alice", "Accounting"},
		{"Bob", "Accounting"},
		{"Cindy", "Engineering"},
		{"Dave", "Engineering"},
		{"Eve", "Engineering"},
	}
	AssertEqual(t, 3, CountBy(employees, func(e Employee) bool {
		return e.department == "Engineering"
	}))
}

func TestDistributionBy(t *testing.T) {
	type Employee struct {
		name       string
		department string
	}
	employees := []Employee{
		{"Alice", "Accounting"},
		{"Bob", "Accounting"},
		{"Cindy", "Engineering"},
		{"Dave", "Engineering"},
		{"Eve", "Engineering"},
	}
	dist := CounterBy(employees, func(e Employee) string {
		return e.department
	})
	AssertMapEqual(t, map[string]int{"Accounting": 2, "Engineering": 3}, dist)
}

func TestUniqBy(t *testing.T) {
	type Employee struct {
		name       string
		department string
	}
	employees := []Employee{
		{"Alice", "Accounting"},
		{"Bob", "Accounting"},
		{"Cindy", "Engineering"},
		{"Dave", "Engineering"},
		{"Eve", "Engineering"},
	}
	uniq := UniqBy(employees, func(e Employee) string {
		return e.department
	})
	expected := []Employee{
		{"Alice", "Accounting"},
		{"Cindy", "Engineering"},
	}
	AssertSliceEqual(t, expected, uniq)
}

func TestUnionBy(t *testing.T) {
	type Employee struct {
		name       string
		department string
	}
	group1 := []Employee{
		{"Alice", "Accounting"},
		{"Bob", "Accounting"},
		{"Cindy", "Engineering"},
	}
	group2 := []Employee{
		{"Alice", "Accounting"},
		{"Cindy", "Engineering"},
		{"Dave", "Engineering"},
		{"Eve", "Engineering"},
	}
	union := UnionBy(func(e Employee) string { return e.name }, group1, group2)
	expected := []Employee{
		{"Alice", "Accounting"},
		{"Bob", "Accounting"},
		{"Cindy", "Engineering"},
		{"Dave", "Engineering"},
		{"Eve", "Engineering"},
	}
	AssertSliceEqual(t, expected, union)
}

func TestIntersectionBy(t *testing.T) {
	type Data struct {
		value int
	}
	data1 := []Data{{1}, {3}, {2}, {4}, {5}}
	data2 := []Data{{2}, {3}, {4}, {5}, {6}}
	data3 := []Data{{5}, {4}, {3}, {2}}
	data4 := []Data{{2}, {3}}
	intersect := IntersectionBy(func(d Data) int { return d.value }, data1, data2, data3, data4)
	expected := []Data{{3}, {2}}
	AssertSliceEqual(t, expected, intersect)

	AssertPanics(t, func() {
		IntersectionBy(func(d Data) int { return d.value }, data1)
	})
}

func TestDiffBy(t *testing.T) {
	type Data struct {
		value int
	}
	data1 := []Data{{1}, {3}, {2}, {4}, {5}, {2}}
	data2 := []Data{{3}, {4}, {5}, {6}}
	data3 := []Data{{5}, {4}, {3}}
	data4 := []Data{{3}, {4}}
	intersect := DifferenceBy(func(d Data) int { return d.value }, data1, data2, data3, data4)
	expected := []Data{{1}, {2}, {2}}
	AssertSliceEqual(t, expected, intersect)
}

func TestChunk(t *testing.T) {
	{
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
		chunks := Chunk(arr, 1)
		expected := [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}}
		for i := range chunks {
			AssertSliceEqual(t, expected[i], chunks[i], strconv.Itoa(i))
		}
	}
	{
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
		chunks := Chunk(arr, 4)
		expected := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
		for i := range chunks {
			AssertSliceEqual(t, expected[i], chunks[i], strconv.Itoa(i))
		}
	}
	{
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
		chunks := Chunk(arr, 3)
		expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}
		for i := range chunks {
			AssertSliceEqual(t, expected[i], chunks[i], strconv.Itoa(i))
		}
	}

	AssertPanics(t, func() {
		Chunk([]int{1, 2, 3}, 0)
	})
}

func TestEqualBy(t *testing.T) {
	{
		a := []int{1, 2, 3, 4, 5}
		b := []rune{'a', 'b', 'c', 'd', 'e'}
		AssertTrue(t, EqualBy(a, b, func(aa int, bb rune) bool {
			return (aa - 1) == int(bb-'a')
		}))
	}
	{
		a := []int{1, 2, 3, 4, 5}
		b := []rune{'a', 'b', 'c', 'd', 'f'}
		AssertFalse(t, EqualBy(a, b, func(aa int, bb rune) bool {
			return (aa - 1) == int(bb-'a')
		}))
	}
	{
		a := []int{1, 2, 3}
		b := []rune{'a', 'b', 'c', 'd', 'f'}
		AssertFalse(t, EqualBy(a, b, func(aa int, bb rune) bool {
			return (aa - 1) == int(bb-'a')
		}))
	}
}
