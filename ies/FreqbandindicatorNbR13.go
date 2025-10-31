package ies

import "rrc/utils"

// FreqBandIndicator-NB-r13 ::= utils.INTEGER (1.. maxFBI2)
type FreqbandindicatorNbR13 struct {
	Value utils.INTEGER `lb:0,ub:maxFBI2`
}
