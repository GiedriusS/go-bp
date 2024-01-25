package bp

import "encoding/binary"

func bitsToMask(bits uint8) uint32 {
	return (1 << bits) - 1
}

type unpacker func([]byte) ([]byte, []uint32)

var unpackerMap = map[uint8]unpacker{
	//1: unpack1,
	//2: unpack2,
	3: unpack3,
	//4: unpack4,
	//5: unpack5,
	//6: unpack6,
	//7: unpack7,
	8:  unpack8,
	10: unpack10,
}

func unpack3(data []byte) ([]byte, []uint32) {
	ret := []uint32{}

	regs := [4]uint32{}

	// Three times (3 * 4 uint32s = 3 * 4 * 4 bytes = 48 bytes).
	{
		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		ret = append(ret, regs[0]&7)
		ret = append(ret, regs[1]&7)
		ret = append(ret, regs[2]&7)
		ret = append(ret, regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3
	}

	// 2 bits remainder, only take 1 bit from the first element.
	{
		remainder := [4]uint32{}
		remainder[0] = regs[0]
		remainder[1] = regs[1]
		remainder[2] = regs[2]
		remainder[3] = regs[3]

		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		ret = append(ret, ret[len(ret)-4]+((regs[0]&1)<<2)|(remainder[0]&3))
		ret = append(ret, ret[len(ret)-4]+((regs[1]&1)<<2)|(remainder[1]&3))
		ret = append(ret, ret[len(ret)-4]+((regs[2]&1)<<2)|(remainder[2]&3))
		ret = append(ret, ret[len(ret)-4]+((regs[3]&1)<<2)|(remainder[3]&3))

		// Push all regs back by one.
		regs[0] >>= 1
		regs[1] >>= 1
		regs[2] >>= 1
		regs[3] >>= 1

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3
	}

	{
		// 1 is left over from the previous one.
		remainder := [4]uint32{}
		remainder[0] = regs[0]
		remainder[1] = regs[1]
		remainder[2] = regs[2]
		remainder[3] = regs[3]

		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		ret = append(ret, ret[len(ret)-4]+((regs[0]&3)<<1)|(remainder[0]&1))
		ret = append(ret, ret[len(ret)-4]+((regs[1]&3)<<1)|(remainder[1]&1))
		ret = append(ret, ret[len(ret)-4]+((regs[2]&3)<<1)|(remainder[2]&1))
		ret = append(ret, ret[len(ret)-4]+((regs[3]&3)<<1)|(remainder[3]&1))

		regs[0] >>= 2
		regs[1] >>= 2
		regs[2] >>= 2
		regs[3] >>= 2

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)

		regs[0] >>= 3
		regs[1] >>= 3
		regs[2] >>= 3
		regs[3] >>= 3

		ret = append(ret, ret[len(ret)-4]+regs[0]&7)
		ret = append(ret, ret[len(ret)-4]+regs[1]&7)
		ret = append(ret, ret[len(ret)-4]+regs[2]&7)
		ret = append(ret, ret[len(ret)-4]+regs[3]&7)
	}

	return data, ret
}

func unpack8(data []byte) ([]byte, []uint32) {
	ret := []uint32{}

	regs := [4]uint32{}

	mask := bitsToMask(8)

	{
		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0])
		ret = append(ret, ret[len(ret)-4]+regs[1])
		ret = append(ret, ret[len(ret)-4]+regs[2])
		ret = append(ret, ret[len(ret)-4]+regs[3])
	}

	{
		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0])
		ret = append(ret, ret[len(ret)-4]+regs[1])
		ret = append(ret, ret[len(ret)-4]+regs[2])
		ret = append(ret, ret[len(ret)-4]+regs[3])
	}

	{
		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8
	}

	{
		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8
	}

	{
		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8
	}

	{
		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8
	}

	{
		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8
	}

	{
		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 8
		regs[1] >>= 8
		regs[2] >>= 8
		regs[3] >>= 8

		ret = append(ret, ret[len(ret)-4]+regs[0])
		ret = append(ret, ret[len(ret)-4]+regs[1])
		ret = append(ret, ret[len(ret)-4]+regs[2])
		ret = append(ret, ret[len(ret)-4]+regs[3])
	}

	return data, ret
}

func unpack10(data []byte) ([]byte, []uint32) {
	ret := []uint32{}
	regs := [4]uint32{}

	// Pulls n bits from the current regs
	// and extracts 10-n from the next regs.
	pulloverBits := func(n uint8) {
		remainderMask := bitsToMask(n)
		leftoverMask := bitsToMask(10 - n)

		remainder := [4]uint32{}
		remainder[0] = regs[0]
		remainder[1] = regs[1]
		remainder[2] = regs[2]
		remainder[3] = regs[3]

		regs[0] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[1] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[2] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]
		regs[3] = binary.LittleEndian.Uint32(data[:4])
		data = data[4:]

		fromPrevious := (remainder[0] & remainderMask)
		cur := ((regs[0] & leftoverMask) << n)
		ret = append(ret, ret[len(ret)-4]+(cur|fromPrevious))
		ret = append(ret, ret[len(ret)-4]+(((regs[1]&leftoverMask)<<n)|(remainder[1]&remainderMask)))
		ret = append(ret, ret[len(ret)-4]+(((regs[2]&leftoverMask)<<n)|(remainder[2]&remainderMask)))
		ret = append(ret, ret[len(ret)-4]+(((regs[3]&leftoverMask)<<n)|(remainder[3]&remainderMask)))

		// Push all regs back by one.
		regs[0] >>= 10 - n
		regs[1] >>= 10 - n
		regs[2] >>= 10 - n
		regs[3] >>= 10 - n
	}

	for i := 0; i < 2; i++ {
		mask := bitsToMask(10)

		{
			regs[0] = binary.LittleEndian.Uint32(data[:4])
			data = data[4:]
			regs[1] = binary.LittleEndian.Uint32(data[:4])
			data = data[4:]
			regs[2] = binary.LittleEndian.Uint32(data[:4])
			data = data[4:]
			regs[3] = binary.LittleEndian.Uint32(data[:4])
			data = data[4:]

			if i > 0 {
				ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
				ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
				ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
				ret = append(ret, ret[len(ret)-4]+regs[3]&mask)
			} else {
				ret = append(ret, regs[0]&mask)
				ret = append(ret, regs[1]&mask)
				ret = append(ret, regs[2]&mask)
				ret = append(ret, regs[3]&mask)
			}

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10

			ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10

			ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10

		}

		{
			pulloverBits(2)

			ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10

			ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10
		}

		{
			pulloverBits(4)

			ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10

			ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10
		}

		{
			pulloverBits(6)

			ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10

			ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10
		}

		{
			pulloverBits(8)

			ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10

			ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10

			ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
			ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

			regs[0] >>= 10
			regs[1] >>= 10
			regs[2] >>= 10
			regs[3] >>= 10
		}
	}

	return data, ret
}
