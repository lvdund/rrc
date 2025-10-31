package ies

import "rrc/utils"

// FreqBandIndicatorEUTRA ::= utils.INTEGER (1..maxBandsEUTRA)
type Freqbandindicatoreutra struct {
	Value utils.INTEGER `lb:0,ub:maxBandsEUTRA`
}
