package ies

import "rrc/utils"

// CellInfoListUTRA-FDD-r9 ::= SEQUENCE OF CellInfoUTRA-FDD-r9
// SIZE (1..maxCellInfoUTRA-r9)
type CellinfolistutraFddR9 struct {
	Value utils.Sequence[CellinfoutraFddR9]
}
