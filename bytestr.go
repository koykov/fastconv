package fastconv

// Fast alloc-free conversion algorithms between byte sequences and strings.

import (
	"reflect"
	"unsafe"
)

// Fast conversion of byte sequence to string.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Shorthand alias of BytesToString().
func B2S(b []byte) string {
	return BytesToString(b)
}

// Fast conversion of string to byte sequence.
func StringToBytes(s string) []byte {
	strh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	var sh reflect.SliceHeader
	sh.Data = strh.Data
	sh.Len = strh.Len
	sh.Cap = strh.Len
	return *(*[]byte)(unsafe.Pointer(&sh))
}

// Shorthand alias of StringToBytes().
func S2B(s string) []byte {
	return StringToBytes(s)
}
