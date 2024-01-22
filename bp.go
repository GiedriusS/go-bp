package bp

import (
	"bytes"
	"fmt"
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
	c := &CompositeDecoder{
		codec1: &nullDecoder{},
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
