# Fast conversion algorithms

## Byte-string-rune conversion

At the moment supports the following conversion:
* `[]byte` -> `string` using `BytesToString()/B2S()`
* `string` -> `[]byte` using `StringToBytes()/S2B()`
* `[]byte` -> `[]rune` using `AppendBytesToRunes()/AppendB2R()`
* `[]rune` -> `[]byte` using `AppendRunesToBytes()/AppendR2B()`
* `string` -> `[]rune` using `AppendStringToRunes()/AppendS2R()`
* `[]rune` -> `string` using `AppendRunesToString()/AppendR2S()`

All functions perform fast conversions using unsafe cast and/or buffer approach.
