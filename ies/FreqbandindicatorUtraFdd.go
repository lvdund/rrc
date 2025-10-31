package ies

import "rrc/utils"

// FreqBandIndicator-UTRA-FDD ::= utils.INTEGER (1..86)
type FreqbandindicatorUtraFdd struct {
	Value utils.INTEGER `lb:0,ub:86`
}
