package ies

import "rrc/utils"

// SK-Counter ::= utils.INTEGER (0..65535)
type SkCounter struct {
	Value utils.INTEGER `lb:0,ub:65535`
}
