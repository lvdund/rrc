package ies

import "rrc/utils"

// MaxMIMO-LayerPreference-r16-reducedMaxMIMO-LayersFR1-r16 ::= SEQUENCE
type MaxmimoLayerpreferenceR16ReducedmaxmimoLayersfr1R16 struct {
	ReducedmimoLayersfr1DlR16 utils.INTEGER `lb:0,ub:8`
	ReducedmimoLayersfr1UlR16 utils.INTEGER `lb:0,ub:4`
}
