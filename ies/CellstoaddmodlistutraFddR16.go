package ies

// CellsToAddModListUTRA-FDD-r16 ::= SEQUENCE OF CellsToAddModUTRA-FDD-r16
// SIZE (1..maxCellMeasUTRA-FDD-r16)
type CellstoaddmodlistutraFddR16 struct {
	Value []CellstoaddmodutraFddR16 `lb:1,ub:maxCellMeasUTRAFddR16`
}
