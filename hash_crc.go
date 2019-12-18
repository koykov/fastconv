package fastconv

const (
	Polynomial int64 = 0xEDB88320

	MaxSlice16    = 16
	MaxSlice8     = 8
	MaxSlice4     = 4
	MaxSlice1     = 1
	MaxSliceNoLut = 0
)

func swap(x uint32) uint32 {
	return (x >> 24) |
		((x >> 8) & 0x0000FF00) |
		((x << 8) & 0x00FF0000) |
		(x << 24)
}

func Crc32Bitwise(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF
	for _, c := range data {
		crc ^= int64(c)
		for j := 0; j < 8; j++ {
			crc = (crc >> 1) ^ (-int64(crc&1) & Polynomial)
		}
	}
	return uint32(crc ^ 0xFFFFFFFF)
}

func Crc32Halfbyte(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF
	for _, c := range data {
		crc = int64(Crc32Lookup16[(uint32(crc)^uint32(c))&0x0f] ^ (uint32(crc) >> 4))
		crc = int64(Crc32Lookup16[(uint32(crc)^(uint32(c)>>4))&0x0f] ^ (uint32(crc) >> 4))
	}
	return uint32(crc ^ 0xFFFFFFFF)
}

func Crc32Byte1(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF
	for _, c := range data {
		crc = (crc >> 8) ^ int64(Crc32Lookup[MaxSliceNoLut][(crc&0xFF)^int64(c)])
	}
	return uint32(crc ^ 0xFFFFFFFF)
}

func Crc32Byte1Tableless(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF
	for _, c := range data {
		s := uint8(crc) ^ uint8(c)
		low := int64((s ^ (s << 6)) & 0xFF)
		a := low * ((1 << 23) + (1 << 14) + (1 << 2))
		crc = (crc >> 8) ^
			(low * ((1 << 24) + (1 << 16) + (1 << 8))) ^
			a ^
			(a >> 1) ^
			(low * ((1 << 20) + (1 << 12))) ^
			(low << 19) ^
			(low << 17) ^
			(low >> 2)
	}
	return uint32(crc ^ 0xFFFFFFFF)
}

func Crc32Byte1Tableless2(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF
	for _, b := range data {
		crc = crc ^ int64(b)
		c := (((crc << 31) >> 31) & ((Polynomial >> 7) ^ (Polynomial >> 1))) ^
			(((crc << 30) >> 31) & ((Polynomial >> 6) ^ Polynomial)) ^
			(((crc << 29) >> 31) & (Polynomial >> 5)) ^
			(((crc << 28) >> 31) & (Polynomial >> 4)) ^
			(((crc << 27) >> 31) & (Polynomial >> 3)) ^
			(((crc << 26) >> 31) & (Polynomial >> 2)) ^
			(((crc << 25) >> 31) & (Polynomial >> 1)) ^
			(((crc << 24) >> 31) & Polynomial)
		crc = int64((crc >> 8) ^ c)
	}
	return uint32(crc ^ 0xFFFFFFFF)
}
