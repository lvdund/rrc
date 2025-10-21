package ies

import "rrc/utils"

// SubframeBitmapSL-r14 ::= CHOICE
type SubframebitmapslR14 interface {
	isSubframebitmapslR14()
}

type SubframebitmapslR14Bs10R14 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR14Bs10R14) isSubframebitmapslR14() {}

type SubframebitmapslR14Bs16R14 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR14Bs16R14) isSubframebitmapslR14() {}

type SubframebitmapslR14Bs20R14 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR14Bs20R14) isSubframebitmapslR14() {}

type SubframebitmapslR14Bs30R14 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR14Bs30R14) isSubframebitmapslR14() {}

type SubframebitmapslR14Bs40R14 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR14Bs40R14) isSubframebitmapslR14() {}

type SubframebitmapslR14Bs50R14 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR14Bs50R14) isSubframebitmapslR14() {}

type SubframebitmapslR14Bs60R14 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR14Bs60R14) isSubframebitmapslR14() {}

type SubframebitmapslR14Bs100R14 struct {
	Value utils.BITSTRING
}

func (*SubframebitmapslR14Bs100R14) isSubframebitmapslR14() {}
