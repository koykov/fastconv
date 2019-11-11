package fastconv

import (
	"bytes"
	"testing"
)

func BenchmarkBytesToString(b *testing.B) {
	p := []byte("foobar")
	s := "foobar"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := BytesToString(p)
		if s != c {
			b.Error(s, "not equal to", s)
		}
	}
}

func BenchmarkBytesToStringNative(b *testing.B) {
	b.ReportAllocs()
	p := []byte("foobar")
	s := "foobar"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := string(p)
		if s != c {
			b.Error(s, "not equal to", s)
		}
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	p := []byte("foobar")
	s := "foobar"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := StringToBytes(s)
		if !bytes.Equal(c, p) {
			b.Error(c, "not equal to", p)
		}
	}
}

func BenchmarkStringToBytesNative(b *testing.B) {
	p := []byte("foobar")
	s := "foobar"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := []byte(s)
		if !bytes.Equal(c, p) {
			b.Error(c, "not equal to", p)
		}
	}
}
