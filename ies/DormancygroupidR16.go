package ies

import "rrc/utils"

// DormancyGroupID-r16 ::= utils.INTEGER (0..4)
type DormancygroupidR16 struct {
	Value utils.INTEGER `lb:0,ub:4`
}
