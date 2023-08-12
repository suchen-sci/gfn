# gfn

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/suchen-sci/gfn/blob/main/LICENSE)

`gfn` is a Golang library that leverages generics to provide various methods, including common functional programming techniques such as `Map`, `Reduce`, and `Filter`, along with other utilities like `Contains`, `Keys`, etc.

1. No `reflect`. 
2. No third-party packages. 
3. Time complexity of `O(n)`.

> If you're interested in contributing to the project, please spare a moment to read [here](./doc/README.md).

## Why this lib?
My friend once complained to me that `Go` is too simple, apart from the essentials, there's hardly anything else. Want to reverse an array? Not available! As a die-hard Gopher, I decided to do something, and hence this library was born. The idea of this library is very simple, it aims to port as many small utilities from other languages to `Go` as possible. The implementation is highly influenced by`Python`, `Ruby`, `JavaScript` and `Lodash`.

On 2023/08/08, with the release of 1.21, `Go` introduced new built-in functions: `min` and `max`, and new standard libraries: `slices` and `maps`. I'm thrilled about these additions. Although 30% of this library can now be replaced by the new built-ins and standard libraries, I've chosen to release this package regardless (still got 70%, right?). I genuinely enjoyed crafting this package, it was really fun.



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
