package gfn_test

import (
	"fmt"
	"math"
	"runtime/debug"
	"strings"
	"testing"

	. "github.com/suchen-sci/gfn"
)

func fail(t *testing.T, failMsg string, tags ...string) {
	t.Helper()
	if len(tags) == 0 {
		t.Error(failMsg)
	} else {
		t.Errorf("%s, tags: %s", failMsg, strings.Join(tags, ", "))
	}
}

func AssertEqual[T comparable](t *testing.T, expected T, actual T, tags ...string) {
	t.Helper()
	if expected != actual {
		fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), tags...)
	}
}

func AssertFloatEqual[T Float](t *testing.T, expected T, actual T, tags ...string) {
	t.Helper()
	if math.Abs(float64(expected-actual)) > 0.0001 {
		fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), tags...)
	}
}

func AssertSliceEqual[T comparable](t *testing.T, expected []T, actual []T, tags ...string) {
	t.Helper()
	if expected == nil || actual == nil {
		if expected == nil && actual == nil {
			return
		}
		fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), tags...)
		return
	}

	if len(expected) != len(actual) {
		fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), tags...)
		return
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), tags...)
			return
		}
	}
}

func AssertMapEqual[T comparable, V comparable](t *testing.T, expected map[T]V, actual map[T]V, tags ...string) {
	t.Helper()
	if expected == nil || actual == nil {
		if expected == nil && actual == nil {
			return
		}
		fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), tags...)
		return
	}

	if len(expected) != len(actual) {
		fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), tags...)
		return
	}

	for k, v := range expected {
		if actualV, ok := actual[k]; !ok || v != actualV {
			fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), tags...)
			return
		}
	}
}

func AssertTrue(t *testing.T, actual bool, tags ...string) {
	t.Helper()
	if !actual {
		fail(t, "expected: true, actual: false", tags...)
	}
}

func AssertFalse(t *testing.T, actual bool, tags ...string) {
	t.Helper()
	if actual {
		fail(t, "expected: false, actual: true", tags...)
	}
}

func AssertPanics(t *testing.T, fn func(), tags ...string) {
	t.Helper()
	defer func() {
		t.Helper()
		if r := recover(); r == nil {
			fail(t, "expected: panic, actual: not panic", tags...)
		}
	}()
	fn()
}

func AssertNotPanics(t *testing.T, fn func(), tags ...string) {
	t.Helper()
	defer func() {
		t.Helper()
		if r := recover(); r != nil {
			stack := string(debug.Stack())
			fail(t, fmt.Sprintf("expected: not panic, actual: %v, %v", r, stack), tags...)
		}
	}()
	fn()
}
