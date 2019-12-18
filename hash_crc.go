package fastconv

const (
	Polynomial uint32 = 0xEDB88320

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
			crc = (crc >> 1) ^ (-int64(crc&1) & int64(Polynomial))
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
