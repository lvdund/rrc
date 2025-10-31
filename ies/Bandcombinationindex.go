package ies

import "rrc/utils"

// BandCombinationIndex ::= utils.INTEGER (1..maxBandComb)
type Bandcombinationindex struct {
	Value utils.INTEGER `lb:0,ub:maxBandComb`
}
