package fastconv

import (
	"hash/fnv"
	"testing"
)

func BenchmarkFnv132(b *testing.B) {
	p := []byte("foobar")
	r := uint32(0x31f0b262)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		h := Fnv132(p)
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkFnv132Native(b *testing.B) {
	p := []byte("foobar")
	r := uint32(0x31f0b262)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f := fnv.New32()
		_, _ = f.Write(p)
		h := f.Sum32()
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkFnv1a32(b *testing.B) {
	p := []byte("foobar")
	r := uint32(0xbf9cf968)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		h := Fnv1a32(p)
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkFnv1a32Native(b *testing.B) {
	p := []byte("foobar")
	r := uint32(0xbf9cf968)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f := fnv.New32a()
		_, _ = f.Write(p)
		h := f.Sum32()
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkFnv164(b *testing.B) {
	p := []byte("foobar")
	r := uint64(0x340d8765a4dda9c2)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		h := Fnv164(p)
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkFnv164Native(b *testing.B) {
	p := []byte("foobar")
	r := uint64(0x340d8765a4dda9c2)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f := fnv.New64()
		_, _ = f.Write(p)
		h := f.Sum64()
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkFnv1a64(b *testing.B) {
	p := []byte("foobar")
	r := uint64(0x85944171f73967e8)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		h := Fnv1a64(p)
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkFnv1a64Native(b *testing.B) {
	p := []byte("foobar")
	r := uint64(0x85944171f73967e8)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f := fnv.New64a()
		_, _ = f.Write(p)
		h := f.Sum64()
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}
