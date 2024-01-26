package main

func bitsToMask(bits uint8) uint32 {
	return (1 << bits) - 1
}
