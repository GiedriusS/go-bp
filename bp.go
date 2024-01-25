package bp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const maxVariableByteLen32 = 5

// VariableByte is like Varint but inverted.
func VariableByte32(buf []byte) (uint32, int) {
	var x uint32
	var s uint
	for i, b := range buf {
		if i == maxVariableByteLen32 {
			// Catch byte reads past maxVariableByteLen32.
			// See issue https://golang.org/issues/41185
			return 0, -(i + 1) // overflow
		}
		x |= uint32(b&0x7f) << s
		s += 7

		// The end!
		if b >= 0x80 {
			if i == maxVariableByteLen32-1 && b > 1 {
				return 0, -(i + 1) // overflow
			}
			return x, i + 1
		}

	}
	return 0, 0
}

type decompressS4BP128D4 struct {
	data []byte

	codec2 bool

	cur uint32
	e   error
}

func NewDecompressS4BP128D4(data []byte) *decompressS4BP128D4 {
	d := &decompressS4BP128D4{data: data}

	// If length of first codec is 0
	// then we start with varint. It is D=1.
	if bytes.HasPrefix(data, []byte{0x00, 0x00, 0x00, 0x00}) {
		d.codec2 = true
		d.data = d.data[4:]
	}
	return d
}

func (d *decompressS4BP128D4) Err() error {
	return d.e
}

func (d *decompressS4BP128D4) Next() bool {
	v, l := VariableByte32(d.data)
	if v == 0 && l <= 0 {
		if l == 0 {
			d.e = fmt.Errorf("buffer too small")
		}
		if l < 0 {
			d.e = fmt.Errorf("value larger than 64 bits")
		}
		return false
	}

	d.cur = d.cur + uint32(v)
	d.data = d.data[l:]

	return true
}

func (d *decompressS4BP128D4) At() uint32 {
	return d.cur
}

func DecompressUnder128(data []byte) []uint32 {
	s4decoder := NewS4BP128D4Decoder(data)
	c := &CompositeDecoder{
		codec1: s4decoder,
		codec2: NewDecompressS4BP128D4(data),
	}
	out := make([]uint32, 0)

	for c.Next() {
		out = append(out, c.At())
	}
	return out
}

type Iterator interface {
	Next() bool
	At() uint32
	Err() error
}

type nullDecoder struct {
}

func (n *nullDecoder) At() uint32 {
	return 0
}

func (n *nullDecoder) Next() bool {
	return false
}

func (n *nullDecoder) Err() error {
	return nil
}

type CompositeDecoder struct {
	codec1, codec2 Iterator
	useCodec2      bool
}

func (d *CompositeDecoder) Next() bool {
	if d.useCodec2 {
		return d.codec2.Next()
	}
	n := d.codec1.Next()
	if !n {
		d.useCodec2 = true
		return d.Next()
	}
	return n
}

func (d *CompositeDecoder) At() uint32 {
	if d.useCodec2 {
		return d.codec2.At()
	}
	return d.codec1.At()
}

func (d *CompositeDecoder) Err() error {
	if d.useCodec2 {
		return d.codec2.Err()
	}
	return d.codec1.Err()
}

type s4BP128D4Decoder struct {
	data []byte
}

type bitsReader struct {
	data     []byte
	usedBits uint8
}

func bitsToMask(bits uint8) uint8 {
	return (1 << bits) - 1
}

func (b *bitsReader) ReadBits(n uint8) (uint32, error) {
	if len(b.data) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	var ret uint32
	if n+b.usedBits > 8 {
		// Switch to the other byte. Use the last bytes.
		ret = uint32(b.data[0] & bitsToMask(n))
		b.usedBits = b.usedBits + n - 8
		b.data = b.data[1:]
		if len(b.data) == 0 {
			return 0, io.ErrUnexpectedEOF
		}
		ret = (ret << uint32(b.usedBits)) | uint32(b.data[0]&bitsToMask(b.usedBits))
		b.data[0] = b.data[0] >> b.usedBits
	} else {
		ret = uint32(b.data[0] & bitsToMask(n))
		b.usedBits += n
		b.data[0] = b.data[0] >> n
	}

	return ret, nil
}

func NewS4BP128D4Decoder(data []byte) *s4BP128D4Decoder {
	// Length refers to number of items coded with the 1st codec.
	length := binary.LittleEndian.Uint32(data[:])

	fmt.Println("number of elements coded with 1st codec", length)
	data = data[4:]

	blockSizes := make([]uint8, 16)
	// Leftovers.
	for i := 0; i < 4; i++ {
		bs := uint8(data[0])
		blockSizes[3+4*i] = bs
		bs = uint8(data[1])
		blockSizes[2+4*i] = bs
		bs = uint8(data[2])
		blockSizes[1+4*i] = bs

		bs = uint8(data[3])
		blockSizes[0+4*i] = bs

		// if bs == 0 then there is nothing there.
		data = data[4:]
	}
	fmt.Println(blockSizes)

	// Now we have bit packed stuff.
	fmt.Println("We are left with", len(data))

	// Dispatch through an array.

	// Create mask.
	// Read 3 bits at a time.
	// AND + store.
	// Shift, AND, store.

	// 1 2 3 4    5 6 7 8
	//            4 4 4 4

	// There should be 32 items (4 bytes each -> uint32).
	// I expect to only see 4 :/

	// Load 4x uint32 LE.
	// 48 bytes = 12 uint32. (12*4)
	//var oldRegs [4]uint32
	unpacked := unpack3(data)
	fmt.Println("Unpacked = ", unpacked, len(unpacked))

	// 16 bytes of bit widths.
	return &s4BP128D4Decoder{data: data}
}

func (d *s4BP128D4Decoder) Next() bool {
	return false
}

func (d *s4BP128D4Decoder) At() uint32 {
	return 0
}

func (d *s4BP128D4Decoder) Err() error {
	return nil
}
