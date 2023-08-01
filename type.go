package gfn

/*
Other basic types:
- bool
- string
- byte, alias for uint8
- rune, alias for int32
*/

// Int contains signed integer types.
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Uint contains unsigned integer types.
type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float contains floating-point types.
type Float interface {
	~float32 | ~float64
}

// Complex contains complex types.
type Complex interface {
	~complex64 | ~complex128
}
