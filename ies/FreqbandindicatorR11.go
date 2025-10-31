package ies

import "rrc/utils"

// FreqBandIndicator-r11 ::= utils.INTEGER (1..maxFBI2)
type FreqbandindicatorR11 struct {
	Value utils.INTEGER `lb:0,ub:maxFBI2`
}
