package ies

// CellsToAddModList ::= SEQUENCE OF CellsToAddMod
// SIZE (1..maxCellMeas)
type Cellstoaddmodlist struct {
	Value []Cellstoaddmod `lb:1,ub:maxCellMeas`
}
