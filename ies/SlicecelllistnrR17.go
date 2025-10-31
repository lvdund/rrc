package ies

// SliceCellListNR-r17 ::= SEQUENCE OF PCI-Range
// SIZE (1..maxCellSlice-r17)
type SlicecelllistnrR17 struct {
	Value []PciRange `lb:1,ub:maxCellSliceR17`
}
