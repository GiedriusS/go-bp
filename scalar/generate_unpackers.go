package main

import (
	"fmt"
	"os"

	. "github.com/dave/jennifer/jen"
)

func generateUnpackers(n int) Code {
	var usedBits uint8

	initCode := []Code{
		Id("mask").Op(":=").Id("bitsToMask").Call(Lit(n)),
		Id("regs").Index(Lit(0)).Op("=").Call(Qual("encoding/binary", "LittleEndian.Uint32").Call(Id("data").Index(Lit(0), Lit(4)))),
		Id("data").Op("=").Id("data").Index(Lit(4), Empty()),
		Id("regs").Index(Lit(1)).Op("=").Call(Qual("encoding/binary", "LittleEndian.Uint32").Call(Id("data").Index(Lit(0), Lit(4)))),
		Id("data").Op("=").Id("data").Index(Lit(4), Empty()),
		Id("regs").Index(Lit(2)).Op("=").Call(Qual("encoding/binary", "LittleEndian.Uint32").Call(Id("data").Index(Lit(0), Lit(4)))),
		Id("data").Op("=").Id("data").Index(Lit(4), Empty()),
		Id("regs").Index(Lit(3)).Op("=").Call(Qual("encoding/binary", "LittleEndian.Uint32").Call(Id("data").Index(Lit(0), Lit(4)))),
		Id("data").Op("=").Id("data").Index(Lit(4), Empty()),
	}

	for i := 0; i < 128/4; i++ {
		if usedBits == 32 {
			usedBits = 0
		} else if usedBits+uint8(n) > 32 {
			initCode = append(initCode, Id("pulloverBits").Call(Lit((32 - usedBits))))
			usedBits = usedBits + uint8(n) - 32
		}

		if i == 0 {
			initCode = append(initCode, Id("ret").Op("=").Append(Id("ret"), Parens(Id("regs").Index(Lit(0)).Op("&").Id("mask"))))
			initCode = append(initCode, Id("ret").Op("=").Append(Id("ret"), Parens(Id("regs").Index(Lit(1)).Op("&").Id("mask"))))
			initCode = append(initCode, Id("ret").Op("=").Append(Id("ret"), Parens(Id("regs").Index(Lit(2)).Op("&").Id("mask"))))
			initCode = append(initCode, Id("ret").Op("=").Append(Id("ret"), Parens(Id("regs").Index(Lit(3)).Op("&").Id("mask"))))

		} else {
			initCode = append(initCode, Id("ret").Op("=").Append(Id("ret"), Id("regs").Index(Len(Id("regs")).Op("-").Lit(4)).Op("+").Parens(Id("regs").Index(Lit(0)).Op("&").Id("mask"))))
			initCode = append(initCode, Id("ret").Op("=").Append(Id("ret"), Id("regs").Index(Len(Id("regs")).Op("-").Lit(4)).Op("+").Parens(Id("regs").Index(Lit(1)).Op("&").Id("mask"))))
			initCode = append(initCode, Id("ret").Op("=").Append(Id("ret"), Id("regs").Index(Len(Id("regs")).Op("-").Lit(4)).Op("+").Parens(Id("regs").Index(Lit(2)).Op("&").Id("mask"))))
			initCode = append(initCode, Id("ret").Op("=").Append(Id("ret"), Id("regs").Index(Len(Id("regs")).Op("-").Lit(4)).Op("+").Parens(Id("regs").Index(Lit(3)).Op("&").Id("mask"))))

		}

		initCode = append(initCode, Id("regs").Index(Lit(0)).Op(">>=").Lit(n))
		initCode = append(initCode, Id("regs").Index(Lit(1)).Op(">>=").Lit(n))
		initCode = append(initCode, Id("regs").Index(Lit(2)).Op(">>=").Lit(n))
		initCode = append(initCode, Id("regs").Index(Lit(3)).Op(">>=").Lit(n))

		usedBits += uint8(n)
	}

	b := Block(
		initCode...,
	)

	return b
}

func writeGenerators(f *File) {
	f.Func().Id("unpackergen3").Params(
		Id("data").Index().Byte(),
	).Parens(List(Index().Byte(), Index().Uint32())).Block(
		Id("regs").Op(":=").Index(Lit(4)).Uint32().Values(),
		Id("ret").Op(":=").Index().Uint32().Values(),
		// pulloverBits takes X bits from the previous registers.
		Id("pulloverBits").Op(":=").Func().Params(Id("n").Uint8()).Block(
			Id("remainderMask").Op(":=").Id("bitsToMask").Call(Id("n")),
			Id("leftoverMask").Op(":=").Id("bitsToMask").Call(Lit(10).Op("-").Id("n")),
			Id("remainder").Op(":=").Index(Lit(4)).Uint32().Values(
				Id("regs").Index(Lit(0)),
				Id("regs").Index(Lit(1)),
				Id("regs").Index(Lit(2)),
				Id("regs").Index(Lit(3)),
			),

			Id("regs").Index(Lit(0)).Op("=").Call(Qual("encoding/binary", "LittleEndian.Uint32").Call(Id("data").Index(Lit(0), Lit(4)))),
			Id("data").Op("=").Id("data").Index(Lit(4), Empty()),
			Id("regs").Index(Lit(0)).Op("=").Call(Qual("encoding/binary", "LittleEndian.Uint32").Call(Id("data").Index(Lit(0), Lit(4)))),
			Id("data").Op("=").Id("data").Index(Lit(4), Empty()),
			Id("regs").Index(Lit(0)).Op("=").Call(Qual("encoding/binary", "LittleEndian.Uint32").Call(Id("data").Index(Lit(0), Lit(4)))),
			Id("data").Op("=").Id("data").Index(Lit(4), Empty()),
			Id("regs").Index(Lit(0)).Op("=").Call(Qual("encoding/binary", "LittleEndian.Uint32").Call(Id("data").Index(Lit(0), Lit(4)))),
			Id("data").Op("=").Id("data").Index(Lit(4), Empty()),

			Id("fromPrevious").Op(":=").Parens(Id("remainder").Index(Lit(0)).Op("&").Id("remainderMask")),
			Id("cur").Op(":=").Parens(Parens(Id("regs").Index(Lit(0)).Op("&").Id("leftoverMask")).Op("<<").Id("n")),
			Id("ret").Op("=").Append(Id("ret"), Id("ret").Index(Len(Id("ret")).Op("-").Lit(4)).Op("+").Parens(Id("cur").Op("|").Id("fromPrevious"))),
			Id("ret").Op("=").Append(Id("ret"), Id("ret").Index(Len(Id("ret")).Op("-").Lit(4)).Op("+").Parens(Parens(Parens(Id("regs").Index(Lit(1)).Op("&").Id("leftoverMask")).Op("<<").Id("n")).Op("|").Parens(Id("remainder").Index(Lit(1)).Op("&").Id("remainderMask")))),
			Id("ret").Op("=").Append(Id("ret"), Id("ret").Index(Len(Id("ret")).Op("-").Lit(4)).Op("+").Parens(Parens(Parens(Id("regs").Index(Lit(2)).Op("&").Id("leftoverMask")).Op("<<").Id("n")).Op("|").Parens(Id("remainder").Index(Lit(2)).Op("&").Id("remainderMask")))),
			Id("ret").Op("=").Append(Id("ret"), Id("ret").Index(Len(Id("ret")).Op("-").Lit(4)).Op("+").Parens(Parens(Parens(Id("regs").Index(Lit(3)).Op("&").Id("leftoverMask")).Op("<<").Id("n")).Op("|").Parens(Id("remainder").Index(Lit(3)).Op("&").Id("remainderMask")))),

			// 3 = how many bits we're pulling.
			Id("regs").Index(Lit(0)).Op(">>=").Lit(3).Op("-").Id("n"),
			Id("regs").Index(Lit(1)).Op(">>=").Lit(3).Op("-").Id("n"),
			Id("regs").Index(Lit(2)).Op(">>=").Lit(3).Op("-").Id("n"),
			Id("regs").Index(Lit(3)).Op(">>=").Lit(3).Op("-").Id("n"),
		),
		// 128 integers.
		// Depending on the integer width, we might need more switching over operations.

		// We need to extract 128 integers.
		// Extract 3 bits over and over.
		// Pull over 3 bits when we need to.

		generateUnpackers(3),
		Return(Id("data"), Id("ret")),
	)
}

func main() {
	f := NewFile("main")

	outf, err := os.Create("../unpackers.go")
	if err != nil {
		panic(err)
	}
	writeGenerators(f)
	if err := f.Render(outf); err != nil {
		panic(err)
	}
	outf.Close()
	fmt.Printf("%#v\n", f)
}
