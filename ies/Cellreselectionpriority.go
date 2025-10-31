package ies

import "rrc/utils"

// CellReselectionPriority ::= utils.INTEGER (0..7)
type Cellreselectionpriority struct {
	Value utils.INTEGER `lb:0,ub:7`
}
