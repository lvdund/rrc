package ies

import "rrc/utils"

// MaxMIMO-LayersDCI-0-2-r16 ::= utils.INTEGER (1..4)
type MaxmimoLayersdci02R16 struct {
	Value utils.INTEGER `lb:0,ub:4`
}
