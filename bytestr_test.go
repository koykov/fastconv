package fastconv

import (
	"bytes"
	"testing"
)

func BenchmarkBytesToString(b *testing.B) {
	b.Run("fastconv", func(b *testing.B) {
		p := []byte("foobar")
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c := BytesToString(p)
			if c != "foobar" {
				b.Errorf("BytesToString mismatch, need %s, got %s", "foobar", c)
			}
		}
	})
	b.Run("native", func(b *testing.B) {
		p := []byte("foobar")
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c := string(p)
			if c != "foobar" {
				b.Errorf("string(bytes) mismatch, need %s, got %s", "foobar", c)
			}
		}
	})
}

func BenchmarkStringToBytes(b *testing.B) {
	b.Run("fastconv", func(b *testing.B) {
		p := []byte("foobar")
		s := "foobar"
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c := StringToBytes(s)
			if !bytes.Equal(c, p) {
				b.Errorf("StringToBytes mismatch, need %s, got %s", "foobar", c)
			}
		}
	})
	b.Run("native", func(b *testing.B) {
		p := []byte("foobar")
		s := "foobar"
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c := []byte(s)
			if !bytes.Equal(c, p) {
				b.Errorf("StringToBytes mismatch, need %s, got %s", "foobar", c)
			}
		}
	})
}
