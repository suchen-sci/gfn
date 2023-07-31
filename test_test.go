package gfn_test

import (
	"fmt"
	"strings"
	"testing"
)

func fail(t *testing.T, failMsg string, msg ...string) {
	t.Helper()
	if len(msg) == 0 {
		t.Error(failMsg)
	} else {
		t.Errorf("%s, msg: %s", failMsg, strings.Join(msg, " "))
	}
}

func AssertEqual[T comparable](t *testing.T, expected T, actual T, msg ...string) {
	t.Helper()
	if expected != actual {
		fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), msg...)
	}
}

func AssertSliceEqual[T comparable](t *testing.T, expected []T, actual []T, msg ...string) {
	t.Helper()
	if expected == nil || actual == nil {
		if expected == nil && actual == nil {
			return
		}
		fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), msg...)
		return
	}

	if len(expected) != len(actual) {
		fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), msg...)
		return
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			fail(t, fmt.Sprintf("expected: %v, actual: %v", expected, actual), msg...)
			return
		}
	}
}

func AssertTrue(t *testing.T, actual bool, msg ...string) {
	t.Helper()
	if !actual {
		fail(t, "expected: true, actual: false", msg...)
	}
}

func AssertFalse(t *testing.T, actual bool, msg ...string) {
	t.Helper()
	if actual {
		fail(t, "expected: false, actual: true", msg...)
	}
}
