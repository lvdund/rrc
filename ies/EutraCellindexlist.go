package ies

// EUTRA-CellIndexList ::= SEQUENCE OF EUTRA-CellIndex
// SIZE (1..maxCellMeasEUTRA)
type EutraCellindexlist struct {
	Value []EutraCellindex `lb:1,ub:maxCellMeasEUTRA`
}
