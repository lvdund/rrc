package ies

import "rrc/utils"

// MaxMIMO-LayerPreferenceFR2-2-r17-reducedMaxMIMO-LayersFR2-2-r17 ::= SEQUENCE
type MaxmimoLayerpreferencefr22R17ReducedmaxmimoLayersfr22R17 struct {
	ReducedmimoLayersfr22DlR17 utils.INTEGER `lb:0,ub:8`
	ReducedmimoLayersfr22UlR17 utils.INTEGER `lb:0,ub:4`
}
