package ies

import "rrc/utils"

// PositionStateVector-r17 ::= utils.INTEGER (-33554432..33554431)
type PositionstatevectorR17 struct {
	Value utils.INTEGER `lb:0,ub:33554431`
}
