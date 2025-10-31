package ies

import "rrc/utils"

// CO-Duration-r16 ::= utils.INTEGER (0..1120)
type CoDurationR16 struct {
	Value utils.INTEGER `lb:0,ub:1120`
}
