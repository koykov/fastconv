package fastconv

// Fast alloc-free conversion algorithms between byte sequences and strings.

import (
	"reflect"
	"unsafe"
)

// BytesToString makes fast conversion of byte array to string.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// B2S is a shorthand alias of BytesToString().
func B2S(b []byte) string {
	return BytesToString(b)
}

// StringToBytes makes fast conversion of string to byte sequence.
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	var h reflect.SliceHeader
	h.Data = sh.Data
	h.Len = sh.Len
	h.Cap = sh.Len
	return *(*[]byte)(unsafe.Pointer(&h))
}

// S2B is a shorthand alias of StringToBytes().
func S2B(s string) []byte {
	return StringToBytes(s)
}
