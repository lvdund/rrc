package ies

// ValidityCellList ::= SEQUENCE OF PCI-Range
// SIZE (1.. maxCellMeasIdle-r16)
type Validitycelllist struct {
	Value []PciRange `lb:1,ub:maxCellMeasIdleR16`
}
