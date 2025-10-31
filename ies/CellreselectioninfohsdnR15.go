package ies

import "rrc/utils"

// CellReselectionInfoHSDN-r15 ::= SEQUENCE
type CellreselectioninfohsdnR15 struct {
	CellequivalentsizeR15 utils.INTEGER `lb:0,ub:16`
}
