package ies

import "rrc/utils"

// PhysCellIdRangeUTRA-FDDList-r9 ::= SEQUENCE OF PhysCellIdRangeUTRA-FDD-r9
// SIZE (1..maxPhysCellIdRange-r9)
type PhyscellidrangeutraFddlistR9 struct {
	Value utils.Sequence[PhyscellidrangeutraFddR9]
}
