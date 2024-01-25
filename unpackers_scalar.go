package bp

import "encoding/binary"

func unpack3(data []byte) []uint32 {
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

	return ret
}
