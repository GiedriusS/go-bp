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
	9:  unpack9,
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

func unpack9(data []byte) ([]byte, []uint32) {
	ret := []uint32{}

	regs := [4]uint32{}

	mask := bitsToMask(9)

	pulloverBits := func(n uint8) {
		remainderMask := bitsToMask(n)
		leftoverMask := bitsToMask(9 - n)

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
		regs[0] >>= 9 - n
		regs[1] >>= 9 - n
		regs[2] >>= 9 - n
		regs[3] >>= 9 - n
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

		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9

	}

	{
		pulloverBits(5)

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9
	}

	{
		pulloverBits(1)

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9
	}

	{
		pulloverBits(8)

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9
	}

	{
		pulloverBits(4)

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9

		ret = append(ret, ret[len(ret)-4]+regs[0]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[1]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[2]&mask)
		ret = append(ret, ret[len(ret)-4]+regs[3]&mask)

		regs[0] >>= 9
		regs[1] >>= 9
		regs[2] >>= 9
		regs[3] >>= 9
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

func unpacker3gen(data []byte) ([]byte, []uint32) {
	regs := [4]uint32{}
	ret := []uint32{}
	pulloverBits := func(n uint8) {
		remainderMask := bitsToMask(n)
		leftoverMask := bitsToMask(10 - n)
		remainder := [4]uint32{regs[0], regs[1], regs[2], regs[3]}
		regs[0] = (binary.LittleEndian.Uint32(data[0:4]))
		data = data[4:]
		regs[0] = (binary.LittleEndian.Uint32(data[0:4]))
		data = data[4:]
		regs[0] = (binary.LittleEndian.Uint32(data[0:4]))
		data = data[4:]
		regs[0] = (binary.LittleEndian.Uint32(data[0:4]))
		data = data[4:]
		fromPrevious := (remainder[0] & remainderMask)
		cur := ((regs[0] & leftoverMask) << n)
		ret = append(ret, ret[len(ret)-4]+(cur|fromPrevious))
		ret = append(ret, ret[len(ret)-4]+(((regs[1]&leftoverMask)<<n)|(remainder[1]&remainderMask)))
		ret = append(ret, ret[len(ret)-4]+(((regs[2]&leftoverMask)<<n)|(remainder[2]&remainderMask)))
		ret = append(ret, ret[len(ret)-4]+(((regs[3]&leftoverMask)<<n)|(remainder[3]&remainderMask)))
		regs[0] >>= 3 - n
		regs[1] >>= 3 - n
		regs[2] >>= 3 - n
		regs[3] >>= 3 - n
	}
	{
		mask := bitsToMask(3)
		regs[0] = (binary.LittleEndian.Uint32(data[0:4]))
		data = data[4:]
		regs[1] = (binary.LittleEndian.Uint32(data[0:4]))
		data = data[4:]
		regs[2] = (binary.LittleEndian.Uint32(data[0:4]))
		data = data[4:]
		regs[3] = (binary.LittleEndian.Uint32(data[0:4]))
		data = data[4:]
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x1))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x4))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x7))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xa))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xd))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x10))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x13))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x16))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x19))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x1c))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x1f))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x22))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x25))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x28))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x2b))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x2e))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x31))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x34))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x37))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x3a))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x3d))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x40))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x43))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x46))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x49))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x4c))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x4f))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x52))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x55))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x58))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x5b))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x5e))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x61))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x64))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x67))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x6a))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x6d))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x70))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x73))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x76))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x79))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x7c))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x7f))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x82))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x85))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x88))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x8b))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x8e))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x91))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x94))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x97))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x9a))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x9d))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xa0))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xa3))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xa6))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xa9))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xac))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xaf))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xb2))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xb5))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xb8))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xbb))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xbe))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xc1))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xc4))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xc7))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xca))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xcd))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xd0))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xd3))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xd6))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xd9))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xdc))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xdf))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x3))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x6))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x9))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xc))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0xf))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x12))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x15))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x18))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x1b))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x1e))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x21))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x24))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x27))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x2a))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x2d))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x30))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x33))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x36))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x39))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x3c))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x3f))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x42))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x45))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x48))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x4b))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x4e))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x51))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x54))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x57))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x5a))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x5d))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
		pulloverBits(uint8(0x60))
		ret = append(ret, regs[0]&mask)
		ret = append(ret, regs[1]&mask)
		ret = append(ret, regs[2]&mask)
		ret = append(ret, regs[3]&mask)
	}
	return data, ret
}
