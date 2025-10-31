package ies

// PCI-List ::= SEQUENCE OF PhysCellId
// SIZE (1..maxNrofCellMeas)
type PciList struct {
	Value []Physcellid `lb:1,ub:maxNrofCellMeas`
}
