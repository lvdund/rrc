package ies

// CellList-r15 ::= SEQUENCE OF PhysCellIdRange
// SIZE (1..maxCellMeasIdle-r15)
type CelllistR15 struct {
	Value []Physcellidrange `lb:1,ub:maxCellMeasIdleR15`
}
