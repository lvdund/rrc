package ies

import "rrc/utils"

// FreqBandIndicatorNR-r15 ::= utils.INTEGER (1.. maxFBI-NR-r15)
type FreqbandindicatornrR15 struct {
	Value utils.INTEGER `lb:0,ub:maxFBINrR15`
}
