package ies

import "rrc/utils"

// MUSIM-GapId-r17 ::= utils.INTEGER (0..2)
type MusimGapidR17 struct {
	Value utils.INTEGER `lb:0,ub:2`
}
