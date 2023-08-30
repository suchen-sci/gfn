// Package main is used to generate README.md
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
	err = os.WriteFile(filepath.Join(wd, readmeFile), []byte(readmeStr), 0644)
	if err != nil {
		panic(err)
	}
}

const tocTemplate = `
- [{{ .Name }}](#{{ .Name | toLower }})
{{ range .Fns }}  - [gfn.{{ .Title }}](#gfn{{ .TOC | toLower }})
{{ end }}
`

const contentTemplate = `
## {{ .Name }}
{{ if .Fns }}
{{ range .Fns }}
### gfn.{{ .Title }}
;;;go
{{ .Signature }}
;;;
{{ .Comment }}
{{ if .Example }}
#### Example:
;;;go
{{ .Example }}
;;;
[back to top](#gfn)
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
	if err := tmlp.Execute(&sb, c); err != nil {
		panic(err)
	}
	return strings.TrimSpace(sb.String()) + "\n"
}

func (c *category) content() string {
	sb := strings.Builder{}
	tmpl := template.Must(template.New("function").Parse(strings.ReplaceAll(contentTemplate, ";;;", "```")))
	if err := tmpl.Execute(&sb, c); err != nil {
		panic(err)
	}
	return sb.String()
}

type fnState int

const (
	stateStart fnState = iota
	stateComment
	stateFinish
	stateAbort
)

type function struct {
	Name       string
	Signature  string
	Comment    string
	Example    string
	state      fnState
	deprecated bool
}

func (f *function) Title() string {
	if f.deprecated {
		return f.Name + " (Deprecated)"
	}
	return f.Name
}

func (f *function) TOC() string {
	if f.deprecated {
		return f.Name + "-deprecated"
	}
	return f.Name
}

func (f *function) addComment(line string) {
	if f.state == stateStart {
		f.state = stateComment
		line = strings.TrimPrefix(line, "//")
		line = strings.TrimSpace(line)
		words := strings.SplitN(line, " ", 3)
		if words[0] == "Deprecated:" {
			f.deprecated = true
			f.Name = words[1]
		} else {
			f.Name = words[0]
		}
		f.Comment = line

	} else if f.state == stateComment {
		line = strings.TrimPrefix(line, "//")
		line = strings.TrimSpace(line)
		f.Comment += " " + line

	}
}

func (f *function) addSignature(line string) {
	line = strings.TrimSpace(line)
	nameFirstChar := string(strings.TrimPrefix(line, "func ")[0])
	if nameFirstChar != strings.ToUpper(nameFirstChar) {
		f.state = stateAbort
		return
	}
	line = strings.TrimRight(line, "{")
	f.Signature = line
	f.state = stateFinish
}

func (f *function) finish() bool {
	return f.state == stateFinish
}

func (f *function) abort() bool {
	return f.state == stateAbort
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
