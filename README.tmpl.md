# gfn
`gfn` is a Golang library that leverages generics to provide various methods, including common functional programming techniques such as `Map`, `Reduce`, and `Filter`, along with other utilities like `Contains`, `Keys`, etc.

No `reflect`. No third-party packages. `O(n)`.

## Why this lib?
My friend once complained to me that `Golang` is too simple, apart from the essentials, there's hardly anything else. Want to reverse an array? Not available! As a die-hard Gopher, I decided to do something, and hence this library was born. The idea of this library is very simple, it aims to port as many small utilities from other languages to `Golang` as possible. The implementation mainly refers to the methods from `Python`, `Ruby`, and `JavaScript`. I hope this library can be helpful when using `Golang`.


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
