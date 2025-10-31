package ies

import "rrc/utils"

// BlackCellsToAddMod ::= SEQUENCE
type Blackcellstoaddmod struct {
	Cellindex       utils.INTEGER `lb:0,ub:maxCellMeas`
	Physcellidrange Physcellidrange
}
