package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	wd := os.Getenv("GFNCWD")
	if wd == "" {
		panic("GFNCWD is not set")
	}

	toc := ""
	content := ""

	for i := 0; i < len(categories); i++ {
		cat, err := processCategory(categories[i][0], filepath.Join(wd, categories[i][1]))
		if err != nil {
			panic(err)
		}
		toc += cat.TOC()
		content += cat.Content()
	}

	readmeTemplate := filepath.Join(wd, "readme.template.md")
	readme, err := os.ReadFile(readmeTemplate)
	if err != nil {
		panic(err)
	}
	readmeStr := strings.Replace(string(readme), "{{ TOC }}", toc, 1)
	readmeStr = strings.Replace(string(readmeStr), "{{ CONTENT }}", content, 1)
	os.WriteFile(filepath.Join(wd, "README.md"), []byte(readmeStr), 0644)
}

type STATE int

const (
	STATE_START STATE = iota
	STATE_COMMENT
	STATE_EXAMPLE
	STATE_SIGNATURE
)

var categories = [][2]string{
	{"Array", "array.go"},
	{"Functional", "fp.go"},
	{"Map", "map.go"},
	{"Math", "math.go"},
}

type Category struct {
	Name string
	Fns  []Function
}

func (c *Category) TOC() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("- [%s](#%s)\n", c.Name, strings.ToLower(c.Name)))
	for _, fn := range c.Fns {
		sb.WriteString(fmt.Sprintf("  - [gfn.%s](#gfn%s)\n", fn.Name, strings.ToLower(fn.Name)))
	}
	return sb.String()
}

func (c *Category) Content() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("## %s\n\n", c.Name))
	for _, fn := range c.Fns {
		sb.WriteString(fn.Content())
	}
	return sb.String()
}

type Function struct {
	Name      string
	Comments  []string
	Signature string
	Example   []string
	State     STATE
}

func (f *Function) AddComment(line string) {
	if f.State == STATE_START {
		f.State = STATE_COMMENT
		line = strings.TrimPrefix(line, "//")
		line = strings.TrimSpace(line)
		words := strings.SplitN(line, " ", 2)
		f.Name = words[0]
		f.Comments = append(f.Comments, line)

	} else if f.State == STATE_COMMENT {
		f.Comments = append(f.Comments, line)

	} else if f.State == STATE_EXAMPLE {
		line = strings.TrimPrefix(line, "//")
		line = strings.TrimSpace(line)
		f.Example = append(f.Example, line)
	}
}

func (f *Function) AddSignature(line string) {
	line = strings.TrimSpace(line)
	line = strings.TrimRight(line, "{")
	f.Signature = line
	f.State = STATE_SIGNATURE
}

func (f *Function) AddExample(line string) {
	f.State = STATE_EXAMPLE
}

func (f *Function) Finish() bool {
	return f.State == STATE_SIGNATURE
}

func (f *Function) Content() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("### gfn.%s\n\n", f.Name))
	sb.WriteString(fmt.Sprintf("```go\n%s\n```\n\n", f.Signature))
	sb.WriteString(strings.Join(f.Comments, " ") + "\n\n")
	if len(f.Example) > 0 {
		sb.WriteString(fmt.Sprintf("```go\n%s\n```\n\n", strings.Join(f.Example, "\n")))
	}
	return sb.String()
}

func processCategory(name, filePath string) (*Category, error) {
	lines, err := readFile(filePath)
	if err != nil {
		return nil, err
	}
	cat := Category{}
	cat.Name = name

	fn := Function{}
	for _, line := range lines {
		if strings.HasPrefix(line, "//") && !strings.HasPrefix(line, "// @example") {
			fn.AddComment(line)

		} else if strings.HasPrefix(line, "func") {
			fn.AddSignature(line)

		} else if strings.HasPrefix(line, "// @example") {
			fn.AddExample(line)

		} else {
			fn = Function{}
		}

		if fn.Finish() {
			cat.Fns = append(cat.Fns, fn)
			fn = Function{}
		}
	}
	sort.Slice(cat.Fns, func(i, j int) bool {
		return cat.Fns[i].Name < cat.Fns[j].Name
	})
	return &cat, nil
}

func readFile(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}
