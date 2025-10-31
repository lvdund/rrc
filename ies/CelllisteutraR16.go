package ies

// CellListEUTRA-r16 ::= SEQUENCE OF EUTRA-PhysCellIdRange
// SIZE (1..maxCellMeasIdle-r16)
type CelllisteutraR16 struct {
	Value []EutraPhyscellidrange `lb:1,ub:maxCellMeasIdleR16`
}
