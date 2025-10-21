package ies

import "rrc/utils"

// SubframeBitmapSL-r12 ::= CHOICE
type SubframebitmapslR12 interface {
	isSubframebitmapslR12()
}

type SubframebitmapslR12Bs4R12 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR12Bs4R12) isSubframebitmapslR12() {}

type SubframebitmapslR12Bs8R12 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR12Bs8R12) isSubframebitmapslR12() {}

type SubframebitmapslR12Bs12R12 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR12Bs12R12) isSubframebitmapslR12() {}

type SubframebitmapslR12Bs16R12 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR12Bs16R12) isSubframebitmapslR12() {}

type SubframebitmapslR12Bs30R12 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR12Bs30R12) isSubframebitmapslR12() {}

type SubframebitmapslR12Bs40R12 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR12Bs40R12) isSubframebitmapslR12() {}

type SubframebitmapslR12Bs42R12 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR12Bs42R12) isSubframebitmapslR12() {}
