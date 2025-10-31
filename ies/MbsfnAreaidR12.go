package ies

import "rrc/utils"

// MBSFN-AreaId-r12 ::= utils.INTEGER (0..255)
type MbsfnAreaidR12 struct {
	Value utils.INTEGER `lb:0,ub:255`
}
