package ies

// PhysCellIdRangeUTRA-FDDList-r9 ::= SEQUENCE OF PhysCellIdRangeUTRA-FDD-r9
// SIZE (1..maxPhysCellIdRange-r9)
type PhyscellidrangeutraFddlistR9 struct {
	Value []PhyscellidrangeutraFddR9 `lb:1,ub:maxPhysCellIdRangeR9`
}
