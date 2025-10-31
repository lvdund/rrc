package ies

import "rrc/utils"

// PUCCH-ResourceExt-v1610-interlaceAllocation-r16 ::= SEQUENCE
type PucchResourceextV1610InterlaceallocationR16 struct {
	RbSetindexR16 utils.INTEGER `lb:0,ub:4`
	Interlace0R16 PucchResourceextV1610InterlaceallocationR16Interlace0R16
}
