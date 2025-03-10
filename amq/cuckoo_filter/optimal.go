package cuckoo

func optimalM(n uint64) uint64 {
	pow2 := func(n uint64) uint64 {
		n--
		n |= n >> 1
		n |= n >> 2
		n |= n >> 4
		n |= n >> 8
		n |= n >> 16
		n |= n >> 32
		n++
		return n
	}
	return pow2(n) / bucketsz
}
