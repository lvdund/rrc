package ies

import "rrc/utils"

// SLSSID-r12 ::= utils.INTEGER (0..167)
type SlssidR12 struct {
	Value utils.INTEGER `lb:0,ub:167`
}
