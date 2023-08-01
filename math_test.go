package gfn_test

import (
	"math"
	"testing"

	. "github.com/suchen-sci/gfn"
)

func TestMax(t *testing.T) {
	AssertEqual(t, 3, Max(1, 2, 3), "int")
	AssertEqual(t, int64(3), Max(int64(1), int64(2), int64(3)), "int64")
	AssertEqual(t, int16(10), Max([]int16{1, 5, 9, 10}...), "int16")

	AssertEqual(t, 3.3, Max(1.1, 2.2, 3.3), "float64")
	AssertEqual(t, float32(3.3), Max(float32(1.1), float32(2.2), float32(3.3)), "float32")

	AssertEqual(t, uint(3), Max(uint(1), uint(2), uint(3)), "uint")
	AssertEqual(t, uint64(3), Max(uint64(1), uint64(2), uint64(3)), "uint64")

	AssertEqual(t, "c", Max("ab", "bd", "c"), "string")
	AssertEqual(t, "hello", Max("hello", "Hello"), "string")
	AssertEqual(t, "hellox", Max("hello", "hellox"), "string")

	AssertPanics(t, func() {
		Max([]int{}...)
	})

	// inf
	AssertEqual(t, math.Inf(1), Max(10000000000.99, math.Inf(1)), "float64")
	AssertEqual(t, math.Inf(1), Max(math.Inf(1), 10000000000.99), "float64")

	// math.NaN() is not comparable.
	// math.NaN() > 0 is false, and math.NaN() < 0 is false.
	AssertTrue(t, math.IsNaN(Max(math.NaN(), 1)), "float64")
	AssertEqual(t, 1, Max(1, math.NaN()), "float64")
}

func TestMaxNotNaN(t *testing.T) {
	AssertEqual(t, 3., MaxNotNaN(1., 2., 3., math.NaN()))
	AssertEqual(t, 4., MaxNotNaN(1., 2., 3., 4.))
	AssertEqual(t, 3., MaxNotNaN(math.NaN(), 1., 2., 3.))
	AssertEqual(t, math.Inf(1), MaxNotNaN(math.NaN(), math.NaN(), math.Inf(1), 2., 3.))
	AssertTrue(t, math.IsNaN(MaxNotNaN(math.NaN(), math.NaN(), math.NaN())))

	AssertPanics(t, func() {
		MaxNotNaN([]float64{}...)
	})
}

func TestMin(t *testing.T) {
	AssertEqual(t, 1, Min(1, 2, 3), "int")
	AssertEqual(t, int64(1), Min(int64(1), int64(2), int64(3)), "int64")
	AssertEqual(t, int16(1), Min([]int16{1, 5, 9, 10}...), "int16")

	AssertEqual(t, 1.1, Min(1.1, 2.2, 3.3), "float64")
	AssertEqual(t, float32(1.1), Min(float32(1.1), float32(2.2), float32(3.3)), "float32")

	AssertEqual(t, uint(1), Min(uint(1), uint(2), uint(3)), "uint")
	AssertEqual(t, uint64(1), Min(uint64(1), uint64(2), uint64(3)), "uint64")

	AssertEqual(t, "ab", Min("ab", "bd", "c"), "string")
	AssertEqual(t, "Hello", Min("hello", "Hello"), "string")
	AssertEqual(t, "hello", Min("hello", "hellox"), "string")

	AssertPanics(t, func() {
		Min([]int{}...)
	})

	// inf
	AssertEqual(t, 10000000000.99, Min(10000000000.99, math.Inf(1)), "float64")
	AssertEqual(t, 10000000000.99, Min(math.Inf(1), 10000000000.99), "float64")

	// math.NaN()
	AssertTrue(t, math.IsNaN(Min(math.NaN(), 1)), "float64")
	AssertEqual(t, 1, Min(1, math.NaN()), "float64")
}

func TestMinNotNaN(t *testing.T) {
	AssertEqual(t, 1., MinNotNaN(1., 2., 3., math.NaN()))
	AssertEqual(t, 1., MinNotNaN(math.NaN(), 1., 2., 3.))
	AssertEqual(t, 1., MinNotNaN(math.NaN(), math.NaN(), 1., 2., 3.))
	AssertEqual(t, math.Inf(-1), MinNotNaN(math.NaN(), math.Inf(-1), math.Inf(1), 2., 3.))
	AssertTrue(t, math.IsNaN(MinNotNaN(math.NaN(), math.NaN(), math.NaN())))

	AssertPanics(t, func() {
		MinNotNaN([]float64{}...)
	})
}

func TestSum(t *testing.T) {
	AssertEqual(t, 6, Sum(1, 2, 3), "int")
	AssertEqual(t, int64(6), Sum(int64(1), int64(2), int64(3)), "int64")
	AssertEqual(t, int16(25), Sum([]int16{1, 5, 9, 10}...), "int16")

	AssertFloatEqual(t, 6.6, Sum(1.1, 2.2, 3.3), "float64")
	AssertFloatEqual(t, float32(6.6), Sum(float32(1.1), float32(2.2), float32(3.3)), "float32")

	AssertEqual(t, uint(6), Sum(uint(1), uint(2), uint(3)), "uint")
	AssertEqual(t, uint64(6), Sum(uint64(1), uint64(2), uint64(3)), "uint64")

	AssertEqual(t, "abbdc", Sum("ab", "bd", "c"), "string")

	AssertEqual(t, complex(12, 8), Sum(complex(1, 2), complex(3, 4), complex(8, 2)), "complex128")

	AssertPanics(t, func() {
		Sum([]int{}...)
	})

	// inf
	AssertEqual(t, math.Inf(1), Sum(10000000000.99, math.Inf(1)), "float64")
	AssertEqual(t, math.Inf(-1), Sum(math.Inf(-1), 10000000000.99), "float64")

	// math.NaN()
	AssertTrue(t, math.IsNaN(Sum(math.Inf(-1), math.Inf(1))), "float64")
	AssertTrue(t, math.IsNaN(Sum(1, math.NaN())), "float64")
}

func TestSumNotNaN(t *testing.T) {
	AssertFloatEqual(t, 6.6, SumNotNaN(1.1, 2.2, 3.3, math.NaN()))
	AssertFloatEqual(t, 6.6, SumNotNaN(math.NaN(), 1.1, 2.2, 3.3, math.NaN()))
	AssertFloatEqual(t, math.Inf(1), SumNotNaN(math.NaN(), math.Inf(1), math.Inf(1), 3.3, math.NaN()))
	AssertFloatEqual(t, math.Inf(1), SumNotNaN(math.NaN(), math.Inf(1), math.Inf(1), 3.3))
	AssertTrue(t, math.IsNaN(SumNotNaN(math.NaN(), math.NaN(), math.NaN())))
	AssertTrue(t, math.IsNaN(SumNotNaN(math.Inf(1), math.Inf(-1))))

	AssertPanics(t, func() {
		SumNotNaN([]float64{}...)
	})
}
