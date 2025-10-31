package ies

// CellsToAddModList ::= SEQUENCE OF CellsToAddMod
// SIZE (1..maxNrofCellMeas)
type Cellstoaddmodlist struct {
	Value []Cellstoaddmod `lb:1,ub:maxNrofCellMeas`
}
