package ies

import "rrc/utils"

// PUCCH-ResourceExt-v1610-interlaceAllocation-r16-interlace0-r16 ::= CHOICE
const (
	PucchResourceextV1610InterlaceallocationR16Interlace0R16ChoiceNothing = iota
	PucchResourceextV1610InterlaceallocationR16Interlace0R16ChoiceScs15
	PucchResourceextV1610InterlaceallocationR16Interlace0R16ChoiceScs30
)

type PucchResourceextV1610InterlaceallocationR16Interlace0R16 struct {
	Choice uint64
	Scs15  *utils.INTEGER `lb:0,ub:9`
	Scs30  *utils.INTEGER `lb:0,ub:4`
}
