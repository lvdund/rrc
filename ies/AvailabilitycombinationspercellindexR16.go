package ies

import "rrc/utils"

// AvailabilityCombinationsPerCellIndex-r16 ::= utils.INTEGER (0..maxNrofDUCells-r16)
type AvailabilitycombinationspercellindexR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofDUCellsR16`
}
