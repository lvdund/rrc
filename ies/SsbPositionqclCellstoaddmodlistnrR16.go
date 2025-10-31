package ies

// SSB-PositionQCL-CellsToAddModListNR-r16 ::= SEQUENCE OF SSB-PositionQCL-CellsToAddNR-r16
// SIZE (1..maxCellMeas)
type SsbPositionqclCellstoaddmodlistnrR16 struct {
	Value []SsbPositionqclCellstoaddnrR16 `lb:1,ub:maxCellMeas`
}
