package ies

// CellListNR-r16 ::= SEQUENCE OF PCI-Range
// SIZE (1..maxCellMeasIdle-r16)
type CelllistnrR16 struct {
	Value []PciRange `lb:1,ub:maxCellMeasIdleR16`
}
