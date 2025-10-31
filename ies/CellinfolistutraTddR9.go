package ies

// CellInfoListUTRA-TDD-r9 ::= SEQUENCE OF CellInfoUTRA-TDD-r9
// SIZE (1..maxCellInfoUTRA-r9)
type CellinfolistutraTddR9 struct {
	Value []CellinfoutraTddR9 `lb:1,ub:maxCellInfoUTRAR9`
}
