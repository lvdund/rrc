package ies

// CellsToAddModList-v1610 ::= SEQUENCE OF CellsToAddMod-v1610
// SIZE (1..maxCellMeas)
type CellstoaddmodlistV1610 struct {
	Value []CellstoaddmodV1610 `lb:1,ub:maxCellMeas`
}
