package ies

// CellInfoListUTRA-TDD-r10 ::= SEQUENCE OF CellInfoUTRA-TDD-r10
// SIZE (1..maxCellInfoUTRA-r9)
type CellinfolistutraTddR10 struct {
	Value []CellinfoutraTddR10 `lb:1,ub:maxCellInfoUTRAR9`
}
