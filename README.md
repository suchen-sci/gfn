# gfn
`gfn` is a Golang library that leverages generics to provide various methods, including common functional programming techniques such as `Map`, `Reduce`, and `Filter`, along with other utilities like `Contains`, `Keys`, etc.

- [Installation](#installation)
- [Usage](#usage)
- [Array](#array)
  - [gfn.Contains](#gfncontains)
  - [gfn.Map](#gfnmap)
## Installation
```
go get github.com/suchen-sci/gfn
```

## Usage 
```
import "github.com/suchen-sci/gfn"
```

## Array

### gfn.Contains

```go
func Contains[T comparable](array []T, value T) bool
```

Contains returns true if the array contains the value.

```go
gfn.Contains([]int{1, 2, 3}, 2)               // true
gfn.Contains([]string{"a", "b", "c"}, "b")    // true
gfn.Contains([]bool{true, false, true}, true) // true
gfn.Contains([]float64{1.1, 2.2, 3.3}, 2.2)   // true

type data struct {
    data int
}
gfn.Contains([]data{{data: 1}, {data: 2}}, data{data: 2})  // true

dataA := &data{data: 1}
dataB := &data{data: 1}
gfn.Contains([]*data{dataA, dataB}, dataA)           // true
gfn.Contains([]*data{dataA, dataB}, &data{data: 1})  // false, not same pointer
```

### gfn.Map

```go
func Map[T any, R any](array []T, mapper func(T) R) []R
```

Map returns a new array with the results of calling the mapper function on each element.

```go
// map int to string
Map([]int{1, 2, 3}, func(i int) string {
    return strconv.Itoa(i)
})  // []string{"1", "2", "3"}

type data struct {
    data string
}

// map string to struct
Map([]string{"a", "b", "c"}, func(s string) data {
    return data{data: s}
})  // []data{{data: "a"}, {data: "b"}, {data: "c"}}
```