# Fast conversion algorithms

## Byte-string conversion

`StringToBytes()` is equal to type cast `[]byte(s)`

Benchmark:
```
BenchmarkStringToBytes-8          200000000    7.28 ns/op    0 B/op    0 allocs/op
BenchmarkStringToBytesNative-8    50000000     33.4 ns/op    8 B/op    1 allocs/op
```

`BytesToString()` is equal to type cast `string(b)`

Benchmark:
```
BenchmarkBytesToString-8          200000000    8.44 ns/op    0 B/op    0 allocs/op
BenchmarkBytesToStringNative-8    100000000    13.7 ns/op    0 B/op    0 allocs/op
```

Note than usage of these functions may be unsafe, since original and result has the same underlying pointer.

## FNV-1 calculation

It's a collection of fast and alloc-free FNV-1 hash calculation.

FNV-1 32 benchmark:
```
BenchmarkFnv32-8                  300000000    5.79 ns/op    0 B/op    0 allocs/op
BenchmarkFnv32Native-8            100000000    21.6 ns/op    4 B/op    1 allocs/op
```

FNV-1a 32 benchmark:
```
BenchmarkFnv32a-8                 300000000    5.81 ns/op    0 B/op    0 allocs/op
BenchmarkFnv32aNative-8           100000000    22.0 ns/op    4 B/op    1 allocs/op
```

FNV-1 64 benchmark:
```
BenchmarkFnv64-8                  300000000    5.94 ns/op    0 B/op    0 allocs/op
BenchmarkFnv64Native-8            50000000     24.5 ns/op    8 B/op    1 allocs/op
```

FNV-1a 64 benchmark:
```
BenchmarkFnv64a-8                 200000000    6.51 ns/op    0 B/op    0 allocs/op
BenchmarkFnv64aNative-8           50000000     24.7 ns/op    8 B/op    1 allocs/op
```

Tech note: these function uses loop-rolling to speed-up hash calculation of long strings. On short string it has
no effect or maybe a bit slowly than simple loop.
