package fastconv

// Collection of loop-rolled calculation of FNV hashes.
// Note than loop-rolling is faster than simple loop only on long string (64+ symbols).

const (
	offset32        = uint32(2166136261)
	offset64        = uint64(14695981039346656037)
	prime32         = uint32(16777619)
	prime64         = uint64(1099511628211)
)

// Fast FNV-1 32 hash calculation.
func Fnv132(p []byte) uint32 {
	h := offset32

	for len(p) >= 8 {
		h *= prime32
		h ^= uint32(p[0])
		h *= prime32
		h ^= uint32(p[1])
		h *= prime32
		h ^= uint32(p[2])
		h *= prime32
		h ^= uint32(p[3])
		h *= prime32
		h ^= uint32(p[4])
		h *= prime32
		h ^= uint32(p[5])
		h *= prime32
		h ^= uint32(p[6])
		h *= prime32
		h ^= uint32(p[7])
		p = p[8:]
	}

	if len(p) >= 4 {
		h *= prime32
		h ^= uint32(p[0])
		h *= prime32
		h ^= uint32(p[1])
		h *= prime32
		h ^= uint32(p[2])
		h *= prime32
		h ^= uint32(p[3])
		p = p[4:]
	}

	if len(p) >= 2 {
		h *= prime32
		h ^= uint32(p[0])
		h *= prime32
		h ^= uint32(p[1])
		p = p[2:]
	}

	if len(p) > 0 {
		h *= prime32
		h ^= uint32(p[0])
	}

	return h
}

// Fast FNV-1 32 hash calculation of string.
func Fnv132String(s string) uint32 {
	return Fnv132(StringToBytes(s))
}

// Fast FNV-1a 32 hash calculation.
func Fnv1a32(p []byte) uint32 {
	h := offset32

	for len(p) >= 8 {
		h = (h ^ uint32(p[0])) * prime32
		h = (h ^ uint32(p[1])) * prime32
		h = (h ^ uint32(p[2])) * prime32
		h = (h ^ uint32(p[3])) * prime32
		h = (h ^ uint32(p[4])) * prime32
		h = (h ^ uint32(p[5])) * prime32
		h = (h ^ uint32(p[6])) * prime32
		h = (h ^ uint32(p[7])) * prime32
		p = p[8:]
	}

	if len(p) >= 4 {
		h = (h ^ uint32(p[0])) * prime32
		h = (h ^ uint32(p[1])) * prime32
		h = (h ^ uint32(p[2])) * prime32
		h = (h ^ uint32(p[3])) * prime32
		p = p[4:]
	}

	if len(p) >= 2 {
		h = (h ^ uint32(p[0])) * prime32
		h = (h ^ uint32(p[1])) * prime32
		p = p[2:]
	}

	if len(p) > 0 {
		h = (h ^ uint32(p[0])) * prime32
	}

	return h
}

// Fast FNV-1a 32 hash calculation of string.
func Fnv1a32String(s string) uint32 {
	return Fnv1a32(StringToBytes(s))
}

// Fast FNV-1 64 hash calculation.
func Fnv164(p []byte) uint64 {
	h := offset64

	for len(p) >= 8 {
		h *= prime64
		h ^= uint64(p[0])
		h *= prime64
		h ^= uint64(p[1])
		h *= prime64
		h ^= uint64(p[2])
		h *= prime64
		h ^= uint64(p[3])
		h *= prime64
		h ^= uint64(p[4])
		h *= prime64
		h ^= uint64(p[5])
		h *= prime64
		h ^= uint64(p[6])
		h *= prime64
		h ^= uint64(p[7])
		p = p[8:]
	}

	if len(p) >= 4 {
		h *= prime64
		h ^= uint64(p[0])
		h *= prime64
		h ^= uint64(p[1])
		h *= prime64
		h ^= uint64(p[2])
		h *= prime64
		h ^= uint64(p[3])
		p = p[4:]
	}

	if len(p) >= 2 {
		h *= prime64
		h ^= uint64(p[0])
		h *= prime64
		h ^= uint64(p[1])
		p = p[2:]
	}

	if len(p) > 0 {
		h *= prime64
		h ^= uint64(p[0])
	}

	return h
}

// Fast FNV-1 64 hash calculation of string.
func Fnv164String(s string) uint64 {
	return Fnv164(StringToBytes(s))
}

// Fast FNV-1a 64 hash calculation.
func Fnv1a64(p []byte) uint64 {
	h := offset64

	for len(p) >= 8 {
		h = (h ^ uint64(p[0])) * prime64
		h = (h ^ uint64(p[1])) * prime64
		h = (h ^ uint64(p[2])) * prime64
		h = (h ^ uint64(p[3])) * prime64
		h = (h ^ uint64(p[4])) * prime64
		h = (h ^ uint64(p[5])) * prime64
		h = (h ^ uint64(p[6])) * prime64
		h = (h ^ uint64(p[7])) * prime64
		p = p[8:]
	}

	if len(p) >= 4 {
		h = (h ^ uint64(p[0])) * prime64
		h = (h ^ uint64(p[1])) * prime64
		h = (h ^ uint64(p[2])) * prime64
		h = (h ^ uint64(p[3])) * prime64
		p = p[4:]
	}

	if len(p) >= 2 {
		h = (h ^ uint64(p[0])) * prime64
		h = (h ^ uint64(p[1])) * prime64
		p = p[2:]
	}

	if len(p) > 0 {
		h = (h ^ uint64(p[0])) * prime64
	}

	return h
}

// Fast FNV-1a 64 hash calculation of string.
func Fnv1a64String(s string) uint64 {
	return Fnv1a64(StringToBytes(s))
}
