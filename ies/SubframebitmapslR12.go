package ies

import "rrc/utils"

// SubframeBitmapSL-r12 ::= CHOICE
const (
	SubframebitmapslR12ChoiceNothing = iota
	SubframebitmapslR12ChoiceBs4R12
	SubframebitmapslR12ChoiceBs8R12
	SubframebitmapslR12ChoiceBs12R12
	SubframebitmapslR12ChoiceBs16R12
	SubframebitmapslR12ChoiceBs30R12
	SubframebitmapslR12ChoiceBs40R12
	SubframebitmapslR12ChoiceBs42R12
)

type SubframebitmapslR12 struct {
	Choice  uint64
	Bs4R12  *utils.BITSTRING `lb:4,ub:4`
	Bs8R12  *utils.BITSTRING `lb:8,ub:8`
	Bs12R12 *utils.BITSTRING `lb:12,ub:12`
	Bs16R12 *utils.BITSTRING `lb:16,ub:16`
	Bs30R12 *utils.BITSTRING `lb:30,ub:30`
	Bs40R12 *utils.BITSTRING `lb:40,ub:40`
	Bs42R12 *utils.BITSTRING `lb:42,ub:42`
}
