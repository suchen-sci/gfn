# gfn

[![Go Report Card](https://goreportcard.com/badge/github.com/suchen-sci/gfn?style=flat-square)](https://goreportcard.com/report/github.com/suchen-sci/gfn)
[![Coverage](https://codecov.io/gh/suchen-sci/gfn/branch/main/graph/badge.svg)](https://app.codecov.io/gh/suchen-sci/gfn/tree/main)
[![Tests](https://github.com/suchen-sci/gfn/actions/workflows/test.yml/badge.svg)](https://github.com/suchen-sci/gfn/actions/workflows/test.yml)
[![Releases](https://img.shields.io/github/release/suchen-sci/gfn/all.svg?style=flat-square)](https://github.com/suchen-sci/gfn/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/suchen-sci/gfn/blob/main/LICENSE)

`gfn` is a Golang library that leverages generics to provide various methods, including common functional programming techniques such as `Map`, `Reduce`, and `Filter`, along with other utilities like `Contains`, `Keys`, etc. The idea of this library is very simple, it aims to port as many small utilities from other languages to `Go` as possible. The implementation is highly influenced by`Python`, `Ruby`, `JavaScript` and `Lodash`.

1. No `reflect`. 
2. No third-party packages. 
3. Time complexity of `O(n)`.

- [Documentation](#documentation)
- [Contributing](#contributing)
- [License](#license)

## Documentation

- [Installation](#installation)
- [Usage](#usage)
- [Type](#type)
{{ TOC }}


## Installation
```
go get github.com/suchen-sci/gfn
```

## Usage 
```
import "github.com/suchen-sci/gfn"
```

## Type

```go
/*
byte: alias for uint8
rune: alias for int32
time.Duration: alias for int64
...
*/

type Int interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
    ~float32 | ~float64
}

type Complex interface {
    ~complex64 | ~complex128
}

type Pair[T, U any] struct {
    First  T
    Second U
}
```

{{ CONTENT }}

## Contributing

Format:
```go
/* @example MyFunc
// your examples here
MyFunc(...)
// output: ...
*/

// MyFunc is ...
func MyFunc(args ...) (return values...) {
    // your code here.
}
```

then run following command to update `README.md` (generated from `README.tmpl.md`).

```bash
make doc
```

[back to top](#gfn)

## License
`gfn` is under the MIT license. See the [LICENSE](https://github.com/suchen-sci/gfn/blob/main/LICENSE) file for details.


[back to top](#gfn)
