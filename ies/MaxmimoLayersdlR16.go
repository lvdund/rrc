package ies

import "rrc/utils"

// MaxMIMO-LayersDL-r16 ::= utils.INTEGER (1..8)
type MaxmimoLayersdlR16 struct {
	Value utils.INTEGER `lb:0,ub:8`
}
