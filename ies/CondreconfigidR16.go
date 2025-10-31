package ies

import "rrc/utils"

// CondReconfigId-r16 ::= utils.INTEGER (1.. maxNrofCondCells-r16)
type CondreconfigidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofCondCellsR16`
}
