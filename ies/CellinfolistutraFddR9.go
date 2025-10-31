package ies

// CellInfoListUTRA-FDD-r9 ::= SEQUENCE OF CellInfoUTRA-FDD-r9
// SIZE (1..maxCellInfoUTRA-r9)
type CellinfolistutraFddR9 struct {
	Value []CellinfoutraFddR9 `lb:1,ub:maxCellInfoUTRAR9`
}
