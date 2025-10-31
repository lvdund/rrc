package ies

import "rrc/utils"

// BetaOffsets ::= SEQUENCE
type Betaoffsets struct {
	BetaoffsetackIndex1      *utils.INTEGER `lb:0,ub:31`
	BetaoffsetackIndex2      *utils.INTEGER `lb:0,ub:31`
	BetaoffsetackIndex3      *utils.INTEGER `lb:0,ub:31`
	BetaoffsetcsiPart1Index1 *utils.INTEGER `lb:0,ub:31`
	BetaoffsetcsiPart1Index2 *utils.INTEGER `lb:0,ub:31`
	BetaoffsetcsiPart2Index1 *utils.INTEGER `lb:0,ub:31`
	BetaoffsetcsiPart2Index2 *utils.INTEGER `lb:0,ub:31`
}
