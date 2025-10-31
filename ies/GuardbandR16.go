package ies

import "rrc/utils"

// GuardBand-r16 ::= SEQUENCE
type GuardbandR16 struct {
	StartcrbR16 utils.INTEGER `lb:0,ub:274`
	NrofcrbsR16 utils.INTEGER `lb:0,ub:15`
}
