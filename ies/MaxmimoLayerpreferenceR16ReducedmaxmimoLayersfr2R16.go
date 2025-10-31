package ies

import "rrc/utils"

// MaxMIMO-LayerPreference-r16-reducedMaxMIMO-LayersFR2-r16 ::= SEQUENCE
type MaxmimoLayerpreferenceR16ReducedmaxmimoLayersfr2R16 struct {
	ReducedmimoLayersfr2DlR16 utils.INTEGER `lb:0,ub:8`
	ReducedmimoLayersfr2UlR16 utils.INTEGER `lb:0,ub:4`
}
