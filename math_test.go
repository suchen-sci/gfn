package gfn_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	. "github.com/suchen-sci/gfn"
)

func TestMax(t *testing.T) {
	// ints
	AssertEqual(t, 3, Max(1, 2, 3), "int")
	AssertEqual(t, int8(4), Max(int8(4), int8(2), int8(3)), "int8")
	AssertEqual(t, int16(4), Max(int16(4), int16(2), int16(3)), "int16")
	AssertEqual(t, int32(4), Max(int32(4), int32(2), int32(3)), "int32")
	AssertEqual(t, int64(4), Max(int64(4), int64(2), int64(3)), "int64")

	// uints
	AssertEqual(t, uint(3), Max([]uint{1, 2, 3}...), "uint")
	AssertEqual(t, uint8(3), Max([]uint8{1, 2, 3}...), "uint8")
	AssertEqual(t, uint16(3), Max([]uint16{1, 2, 3}...), "uint16")
	AssertEqual(t, uint32(3), Max([]uint32{1, 2, 3}...), "uint32")
	AssertEqual(t, uint64(3), Max([]uint64{1, 2, 3}...), "uint64")
	AssertEqual(t, uintptr(3), Max([]uintptr{1, 2, 3}...), "uintptr")

	// float32
	AssertEqual(t, float32(3.3), Max([]float32{1.1, 2.2, 3.3}...), "float32")

	// string
	AssertEqual(t, "cd", Max([]string{"abc", "bd", "cd"}...), "string")

	// check ~int
	type MyInt int
	AssertEqual(t, MyInt(3), Max([]MyInt{1, 2, 3}...), "MyInt")

	// check empty array
	AssertPanics(t, func() {
		Max([]int{}...)
	})

	// check array with many elements
	{
		array := make([]int, 100000)
		for i := 0; i < 100000; i++ {
			array[i] = i
		}
		rand.Shuffle(len(array), func(i, j int) {
			array[i], array[j] = array[j], array[i]
		})
		AssertEqual(t, 99999, Max(array...))
	}
}

func TestMaxFloat64(t *testing.T) {
	AssertEqual(t, 2.2, MaxFloat64(math.NaN(), 1, 2.2))
	AssertEqual(t, 2.8, MaxFloat64(1, -1, math.NaN(), 1, 2.8))
	AssertEqual(t, math.Inf(1), MaxFloat64(1, -1, math.NaN(), 1, math.Inf(1)))
	AssertEqual(t, math.Inf(1), MaxFloat64(1, -1, 1, math.Inf(1)))
	AssertEqual(t, 1.9, MaxFloat64(1.9, -1, 1))
	AssertTrue(t, math.IsNaN(MaxFloat64(math.NaN(), math.NaN(), math.NaN())))

	// check empty array
	AssertPanics(t, func() {
		MaxFloat64()
	})

	// check array with many elements
	{
		array := make([]float64, 100000)
		for i := 0; i < 100000; i++ {
			array[i] = float64(i)
		}
		rand.Shuffle(len(array), func(i, j int) {
			array[i], array[j] = array[j], array[i]
		})
		AssertEqual(t, float64(99999), MaxFloat64(array...))
	}
}

func TestMin(t *testing.T) {
	// ints
	AssertEqual(t, 1, Min(1, 2, 3), "int")
	AssertEqual(t, int8(2), Min(int8(4), int8(2), int8(3)), "int8")
	AssertEqual(t, int16(2), Min(int16(4), int16(2), int16(3)), "int16")
	AssertEqual(t, int32(2), Min(int32(4), int32(2), int32(3)), "int32")
	AssertEqual(t, int64(2), Min(int64(4), int64(2), int64(3)), "int64")

	// uints
	AssertEqual(t, uint(1), Min([]uint{1, 2, 3}...), "uint")
	AssertEqual(t, uint8(1), Min([]uint8{1, 2, 3}...), "uint8")
	AssertEqual(t, uint16(1), Min([]uint16{1, 2, 3}...), "uint16")
	AssertEqual(t, uint32(1), Min([]uint32{1, 2, 3}...), "uint32")
	AssertEqual(t, uint64(1), Min([]uint64{1, 2, 3}...), "uint64")
	AssertEqual(t, uintptr(1), Min([]uintptr{1, 2, 3}...), "uintptr")

	// float32
	AssertEqual(t, float32(1.1), Min([]float32{1.1, 2.2, 3.3}...), "float32")

	// string
	AssertEqual(t, "abc", Min([]string{"abc", "bd", "cd"}...), "string")

	// check ~int
	type MyInt int
	AssertEqual(t, MyInt(1), Min([]MyInt{1, 2, 3}...), "MyInt")

	// check empty array
	AssertPanics(t, func() {
		Min([]int{}...)
	})

	// check array with many elements
	{
		array := make([]int, 100000)
		for i := 0; i < 100000; i++ {
			array[i] = i
		}
		rand.Shuffle(len(array), func(i, j int) {
			array[i], array[j] = array[j], array[i]
		})
		AssertEqual(t, 0, Min(array...))
	}
}

func TestMinFloat64(t *testing.T) {
	AssertEqual(t, 1.85, MinFloat64(math.NaN(), 1.85, 2.2))
	AssertEqual(t, -1, MinFloat64(1, -1, math.NaN(), 1, 2.8))
	AssertEqual(t, math.Inf(-1), MinFloat64(1, -1, math.NaN(), 1, math.Inf(-1)))
	AssertEqual(t, -1, MinFloat64(1, -1, 1, math.Inf(1)))
	AssertEqual(t, -1, MinFloat64(1.9, -1, 1))
	AssertTrue(t, math.IsNaN(MinFloat64(math.NaN(), math.NaN(), math.NaN())))

	// check empty array
	AssertPanics(t, func() {
		MinFloat64()
	})

	// check array with many elements
	{
		array := make([]float64, 100000)
		for i := 0; i < 100000; i++ {
			array[i] = float64(i)
		}
		rand.Shuffle(len(array), func(i, j int) {
			array[i], array[j] = array[j], array[i]
		})
		AssertEqual(t, float64(0), MinFloat64(array...))
	}
}

func TestSum(t *testing.T) {
	// ints
	AssertEqual(t, 10, Sum(1, 2, 3, 4), "int")
	AssertEqual(t, int8(10), Sum([]int8{1, 2, 3, 4}...), "int8")
	AssertEqual(t, int16(10), Sum([]int16{1, 2, 3, 4}...), "int16")
	AssertEqual(t, int32(10), Sum([]int32{1, 2, 3, 4}...), "int32")
	AssertEqual(t, int64(10), Sum([]int64{1, 2, 3, 4}...), "int64")

	// uints
	AssertEqual(t, uint(10), Sum([]uint{1, 2, 3, 4}...), "uint")
	AssertEqual(t, uint8(10), Sum([]uint8{1, 2, 3, 4}...), "uint8")
	AssertEqual(t, uint16(10), Sum([]uint16{1, 2, 3, 4}...), "uint16")
	AssertEqual(t, uint32(10), Sum([]uint32{1, 2, 3, 4}...), "uint32")
	AssertEqual(t, uint64(10), Sum([]uint64{1, 2, 3, 4}...), "uint64")
	AssertEqual(t, uintptr(10), Sum([]uintptr{1, 2, 3, 4}...), "uintptr")

	// floats
	AssertFloatEqual(t, float32(10.1), Sum([]float32{1.1, 2.2, 3.3, 3.5}...), "float32")
	AssertFloatEqual(t, 10.1, Sum([]float64{1.1, 2.2, 3.3, 3.5}...), "float64")

	// string
	AssertEqual(t, "abcde", Sum([]string{"ab", "cd", "e"}...), "string")

	// complex
	AssertEqual(t, complex64(10+10i), Sum([]complex64{1 + 1i, 2 + 2i, 3 + 3i, 4 + 4i}...), "complex64")
	AssertEqual(t, complex128(10+10i), Sum([]complex128{1 + 1i, 2 + 2i, 3 + 3i, 4 + 4i}...), "complex128")

	// check ~int
	type MyInt int
	AssertEqual(t, MyInt(10), Sum([]MyInt{1, 2, 3, 4}...), "~int")

	// check empty array
	AssertPanics(t, func() {
		Sum([]int{}...)
	})

	// check array with many elements
	{
		array := make([]int, 100000)
		res := 0
		for i := 0; i < 100000; i++ {
			array[i] = i
			res += i
		}
		rand.Shuffle(len(array), func(i, j int) {
			array[i], array[j] = array[j], array[i]
		})
		AssertEqual(t, res, Sum(array...))
	}

	// check float64 NaN
	AssertTrue(t, math.IsNaN(Sum([]float64{1, 2, 3, math.NaN()}...)))
	AssertTrue(t, math.IsNaN(Sum([]float64{math.NaN(), 2, 3}...)))
	AssertTrue(t, math.IsNaN(Sum([]float64{math.Inf(1), math.Inf(-1), math.Inf(1)}...)))
}

func TestAbs(t *testing.T) {
	AssertEqual(t, 1., Abs(1.))
	AssertEqual(t, 1., Abs(-1.))
	AssertEqual(t, math.Inf(1), Abs(math.Inf(1)))
	AssertEqual(t, math.Inf(1), Abs(math.Inf(-1)))
	AssertTrue(t, math.IsNaN(Abs(math.NaN())))
	AssertTrue(t, math.IsNaN(Abs(-math.NaN())))

	AssertEqual(t, 1, Abs(-1))
	AssertEqual(t, 1, Abs(1))
	AssertEqual(t, int64(100), Abs(int64(-100)))
	AssertEqual(t, int64(100), Abs(int64(100)))
}

func TestDivMod(t *testing.T) {
	for i := 0; i < 1000; i++ {
		a := rand.Intn(10000) + 1
		b := rand.Intn(10000) + 1
		if rand.Intn(1000) < 250 {
			a = -a
		}
		if rand.Intn(1000) < 250 {
			b = -b
		}
		div, mod := DivMod(a, b)
		AssertEqual(t, a/b, div, fmt.Sprintf("%d/%d", a, b))
		AssertEqual(t, a%b, mod, fmt.Sprintf("%d/%d", a, b))
	}

	testCases := []struct {
		a, b, div, mod int
	}{
		{1, 1, 1, 0},
		{1, 2, 0, 1},
		{-10, 3, -3, -1},
		{-11, 3, -3, -2},
		{-12, 3, -4, 0},
		{-13, 3, -4, -1},
		{139, 3, 46, 1},
	}

	for _, tc := range testCases {
		div, mod := DivMod(tc.a, tc.b)
		AssertEqual(t, tc.div, div, fmt.Sprintf("%d/%d", tc.a, tc.b))
		AssertEqual(t, tc.mod, mod, fmt.Sprintf("%d/%d", tc.a, tc.b))
	}

	AssertPanics(t, func() {
		DivMod(1, 0)
	})
}

func TestSumBy(t *testing.T) {
	type Product struct {
		name   string
		amount int
	}
	products := []Product{
		{"apple", 10},
		{"banana", 20},
		{"orange", 30},
	}
	AssertEqual(t, 60, SumBy(products, func(p Product) int {
		return p.amount
	}))

	AssertPanics(t, func() {
		SumBy([]int{}, func(i int) int {
			return i
		})
	})
}

func TestMaxBy(t *testing.T) {
	type Product struct {
		name   string
		amount int
	}
	products := []Product{
		{"apple", 10},
		{"banana", 20},
		{"orange", 30},
	}
	p := MaxBy(products, func(p Product) int {
		return p.amount
	})
	AssertEqual(t, "orange", p.name)

	AssertPanics(t, func() {
		MaxBy([]int{}, func(i int) int {
			return i
		})
	})
}

func TestMinBy(t *testing.T) {
	type Product struct {
		name   string
		amount int
	}
	products := []Product{
		{"banana", 20},
		{"apple", 10},
		{"orange", 30},
	}
	p := MinBy(products, func(p Product) int {
		return p.amount
	})
	AssertEqual(t, "apple", p.name)

	AssertPanics(t, func() {
		MinBy([]int{}, func(i int) int {
			return i
		})
	})
}

func TestMean(t *testing.T) {
	AssertFloatEqual(t, 2.5, Mean([]int{1, 2, 3, 4}...), "int")
	AssertFloatEqual(t, 3.0, Mean([]float64{1.5, 2.5, 3.5, 4.5}...), "int")

	AssertPanics(t, func() {
		Mean([]int{}...)
	})
}

func TestMeanBy(t *testing.T) {
	type Product struct {
		name string
		cost float64
	}
	products := []Product{
		{"apple", 1.5},
		{"banana", 2.5},
		{"orange", 3.5},
		{"lemon", 4.5},
	}
	AssertFloatEqual(t, 3.0, MeanBy(products, func(p Product) float64 {
		return p.cost
	}))

	AssertPanics(t, func() {
		MeanBy([]int{}, func(i int) int {
			return i
		})
	})
}
