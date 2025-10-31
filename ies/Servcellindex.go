package ies

import "rrc/utils"

// ServCellIndex ::= utils.INTEGER (0..maxNrofServingCells-1)
type Servcellindex struct {
	Value utils.INTEGER `lb:0,ub:maxNrofServingCells1`
}
