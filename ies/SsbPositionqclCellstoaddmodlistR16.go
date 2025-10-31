package ies

// SSB-PositionQCL-CellsToAddModList-r16 ::= SEQUENCE OF SSB-PositionQCL-CellsToAddMod-r16
// SIZE (1..maxNrofCellMeas)
type SsbPositionqclCellstoaddmodlistR16 struct {
	Value []SsbPositionqclCellstoaddmodR16 `lb:1,ub:maxNrofCellMeas`
}
