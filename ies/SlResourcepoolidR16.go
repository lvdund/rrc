package ies

import "rrc/utils"

// SL-ResourcePoolID-r16 ::= utils.INTEGER (1..maxNrofPoolID-r16)
type SlResourcepoolidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPoolIDR16`
}
