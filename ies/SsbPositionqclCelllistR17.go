package ies

// SSB-PositionQCL-CellList-r17 ::= SEQUENCE OF SSB-PositionQCL-Cell-r17
// SIZE (1..maxNrofCellMeas)
type SsbPositionqclCelllistR17 struct {
	Value []SsbPositionqclCellR17 `lb:1,ub:maxNrofCellMeas`
}
