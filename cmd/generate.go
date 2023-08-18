package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

// See Contributing part in README.md

var categories = [][2]string{
	{"Functional", "fp.go"},
	{"Math", "math.go"},
	{"Array", "array.go"},
	{"Map", "map.go"},
}

const readmeTemplateFile = "README.tmpl.md"
const readmeFile = "README.md"

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
			panic(fmt.Errorf("process category %s %s failed, %s", categories[i][0], categories[i][1], err.Error()))
		}
		toc += cat.toc()
		content += cat.content()
	}

	readmeTmpl, err := os.ReadFile(filepath.Join(wd, readmeTemplateFile))
	if err != nil {
		panic(err)
	}
	readmeStr := strings.Replace(strings.Replace(string(readmeTmpl), "{{ TOC }}", toc, 1), "{{ CONTENT }}", content, 1)
	os.WriteFile(filepath.Join(wd, readmeFile), []byte(readmeStr), 0644)
}

const tocTemplate = `
- [{{ .Name }}](#{{ .Name | toLower }})
{{ range .Fns }}  - [gfn.{{ .Name }}](#gfn{{ .Name | toLower }})
{{ end }}
`

const contentTemplate = `
## {{ .Name }}
{{ if .Fns }}
{{ range .Fns }}
### gfn.{{ .Name }}
;;;go
{{ .Signature }}
;;;
{{ .Comment }}
{{ if .Example }}
#### Example:
;;;go
{{ .Example }}
;;;
{{ end }}
{{ end }}
{{ end }}
`

type category struct {
	Name string
	Fns  []function
}

func (c *category) toc() string {
	sb := strings.Builder{}
	funcs := template.FuncMap{
		"toLower": strings.ToLower,
	}
	tmlp := template.Must(template.New("toc").Funcs(funcs).Parse(tocTemplate))
	tmlp.Execute(&sb, c)
	return strings.TrimSpace(sb.String()) + "\n"
}

func (c *category) content() string {
	sb := strings.Builder{}
	tmpl := template.Must(template.New("function").Parse(strings.ReplaceAll(contentTemplate, ";;;", "```")))
	tmpl.Execute(&sb, c)
	return sb.String()
}

type fnState int

const (
	state_start fnState = iota
	state_comment
	state_finish
	state_abort
)

type function struct {
	Name      string
	Signature string
	Comment   string
	Example   string
	state     fnState
}

func (f *function) addComment(line string) {
	if f.state == state_start {
		f.state = state_comment
		line = strings.TrimPrefix(line, "//")
		line = strings.TrimSpace(line)
		words := strings.SplitN(line, " ", 2)
		f.Name = words[0]
		f.Comment = line

	} else if f.state == state_comment {
		line = strings.TrimPrefix(line, "//")
		line = strings.TrimSpace(line)
		f.Comment += " " + line

	}
}

func (f *function) addSignature(line string) {
	line = strings.TrimSpace(line)
	nameFirstChar := string(strings.TrimPrefix(line, "func ")[0])
	if nameFirstChar != strings.ToUpper(nameFirstChar) {
		f.state = state_abort
		return
	}
	line = strings.TrimRight(line, "{")
	f.Signature = line
	f.state = state_finish
}

func (f *function) finish() bool {
	return f.state == state_finish
}

func (f *function) abort() bool {
	return f.state == state_abort
}

func processCategory(name, filePath string) (*category, error) {
	lines, err := readFile(filePath)
	if err != nil {
		return nil, err
	}
	cat := category{}
	cat.Name = name

	multiLineComments := map[string][]string{}

	fn := function{}
	i := 0
	for i < len(lines) {
		line := lines[i]
		if strings.HasPrefix(line, "//") {
			fn.addComment(line)

		} else if strings.HasPrefix(line, "func") {
			fn.addSignature(line)

		} else if strings.HasPrefix(line, "/* @example") {
			line = strings.TrimPrefix(line, "/* @example")
			name := strings.TrimSpace(line)
			comments := []string{}
			for j := i + 1; j < len(lines); j++ {
				if strings.HasPrefix(lines[j], "*/") || strings.HasPrefix(lines[j], " */") {
					i = j + 1
					break
				}
				comments = append(comments, strings.ReplaceAll(lines[j], "\t", "    "))
			}
			multiLineComments[name] = comments
		} else {
			fn = function{}
		}

		if fn.finish() {
			cat.Fns = append(cat.Fns, fn)
			fn = function{}
		} else if fn.abort() {
			fn = function{}
		}
		i++
	}

	sort.Slice(cat.Fns, func(i, j int) bool {
		return cat.Fns[i].Name < cat.Fns[j].Name
	})

	for i := range cat.Fns {
		fn = cat.Fns[i]
		if comments, ok := multiLineComments[fn.Name]; ok {
			if fn.Example == "" {
				fn.Example = strings.Join(comments, "\n")
			} else {
				fn.Example = strings.Join(comments, "\n") + "\n\n" + fn.Example
			}
			cat.Fns[i] = fn
		}
	}
	return &cat, nil
}

func readFile(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}
