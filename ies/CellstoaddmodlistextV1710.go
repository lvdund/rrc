package ies

// CellsToAddModListExt-v1710 ::= SEQUENCE OF CellsToAddModExt-v1710
// SIZE (1..maxNrofCellMeas)
type CellstoaddmodlistextV1710 struct {
	Value []CellstoaddmodextV1710 `lb:1,ub:maxNrofCellMeas`
}
