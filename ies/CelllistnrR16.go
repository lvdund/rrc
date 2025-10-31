package ies

// CellListNR-r16 ::= SEQUENCE OF PhysCellIdRangeNR-r16
// SIZE (1..maxCellMeasIdle-r15)
type CelllistnrR16 struct {
	Value []PhyscellidrangenrR16 `lb:1,ub:maxCellMeasIdleR15`
}
