package ies

import "rrc/utils"

// PortIndex8 ::= utils.INTEGER (0..7)
type Portindex8 struct {
	Value utils.INTEGER `lb:0,ub:7`
}
