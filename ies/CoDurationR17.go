package ies

import "rrc/utils"

// CO-Duration-r17 ::= utils.INTEGER (0..4480)
type CoDurationR17 struct {
	Value utils.INTEGER `lb:0,ub:4480`
}
