package ies

import "rrc/utils"

// SubframeBitmapSL-r14 ::= CHOICE
const (
	SubframebitmapslR14ChoiceNothing = iota
	SubframebitmapslR14ChoiceBs10R14
	SubframebitmapslR14ChoiceBs16R14
	SubframebitmapslR14ChoiceBs20R14
	SubframebitmapslR14ChoiceBs30R14
	SubframebitmapslR14ChoiceBs40R14
	SubframebitmapslR14ChoiceBs50R14
	SubframebitmapslR14ChoiceBs60R14
	SubframebitmapslR14ChoiceBs100R14
)

type SubframebitmapslR14 struct {
	Choice   uint64
	Bs10R14  *utils.BITSTRING `lb:10,ub:10`
	Bs16R14  *utils.BITSTRING `lb:16,ub:16`
	Bs20R14  *utils.BITSTRING `lb:20,ub:20`
	Bs30R14  *utils.BITSTRING `lb:30,ub:30`
	Bs40R14  *utils.BITSTRING `lb:40,ub:40`
	Bs50R14  *utils.BITSTRING `lb:50,ub:50`
	Bs60R14  *utils.BITSTRING `lb:60,ub:60`
	Bs100R14 *utils.BITSTRING `lb:100,ub:100`
}
