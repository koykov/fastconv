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
