package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestProcessCategory(t *testing.T) {
	fileData := `

/* @example F1
this is multiline comments for F1.
*/


// F1 is f1.
func F1(a int) int {
	return a
}

// skipFn is a function that should be skipped.
func skipFn(a int) int {
	return a
}
`
	dir, err := os.MkdirTemp("", "test-generate")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	filePath := filepath.Join(dir, "test.go")
	if err := os.WriteFile(filePath, []byte(fileData), 0644); err != nil {
		t.Fatal(err)
	}
	categories = [][2]string{{"Test", "test.go"}}
	cat, err := processCategory(categories[0][0], filePath)
	if err != nil {
		t.Fatal(err)
	}

	toc := `- [Test](#test)
  - [gfn.F1](#gfnf1)
`
	if cat.toc() != toc {
		t.Fatalf("toc not match, expect: %s, got: %s", toc, cat.toc())
	}

	content := ` ## Test


### gfn.F1
;;;go
func F1(a int) int 
;;;
F1 is f1.

#### Example:
;;;go
this is multiline comments for F1.
;;;
`
	expected := strings.TrimSpace(strings.ReplaceAll(content, ";;;", "```"))
	got := strings.TrimSpace(cat.content())
	if expected != got {
		t.Fatalf("content not match, expect: %s, got: %s", expected, got)
	}
}
