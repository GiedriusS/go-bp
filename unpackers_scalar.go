package bp

func bitsToMask(bits uint8) uint32 {
	return (1 << bits) - 1
}

type unpacker func([]byte) ([]byte, []uint32)

var unpackerMap = map[uint8]unpacker{
	1:  unpackergen1,
	2:  unpackergen2,
	3:  unpackergen3,
	4:  unpackergen4,
	5:  unpackergen5,
	6:  unpackergen6,
	7:  unpackergen7,
	8:  unpackergen8,
	9:  unpackergen9,
	10: unpackergen10,
	11: unpackergen11,
	12: unpackergen12,
	13: unpackergen13,
	14: unpackergen14,
	15: unpackergen15,
	16: unpackergen16,
	17: unpackergen17,
	18: unpackergen18,
	19: unpackergen19,
	20: unpackergen20,
	21: unpackergen21,
	22: unpackergen22,
	23: unpackergen23,
	24: unpackergen24,
	25: unpackergen25,
	26: unpackergen26,
	27: unpackergen27,
	28: unpackergen28,
	29: unpackergen29,
	30: unpackergen30,
	31: unpackergen31,
}
