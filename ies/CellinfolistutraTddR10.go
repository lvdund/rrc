package ies

import "rrc/utils"

// CellInfoListUTRA-TDD-r10 ::= SEQUENCE OF CellInfoUTRA-TDD-r10
// SIZE (1..maxCellInfoUTRA-r9)
type CellinfolistutraTddR10 struct {
	Value []CellinfoutraTddR10
}
