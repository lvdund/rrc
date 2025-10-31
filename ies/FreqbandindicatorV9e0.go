package ies

import "rrc/utils"

// FreqBandIndicator-v9e0 ::= utils.INTEGER (maxFBI-Plus1..maxFBI2)
type FreqbandindicatorV9e0 struct {
	Value utils.INTEGER `lb:maxFBIPlus1,ub:maxFBI2`
}
