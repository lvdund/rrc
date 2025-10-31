package ies

import "rrc/utils"

// PortIndex4 ::= utils.INTEGER (0..3)
type Portindex4 struct {
	Value utils.INTEGER `lb:0,ub:3`
}
