package ies

import "rrc/utils"

// PhysCellIdNR-r15 ::= utils.INTEGER (0.. 1007)
type PhyscellidnrR15 struct {
	Value utils.INTEGER `lb:0,ub:1007`
}
