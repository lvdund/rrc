package ies

import "rrc/utils"

// PortIndex2 ::= utils.INTEGER (0..1)
type Portindex2 struct {
	Value utils.INTEGER `lb:0,ub:1`
}
