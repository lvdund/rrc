package ies

// SCPTM-NeighbourCellList-NB-r14 ::= SEQUENCE OF PCI-ARFCN-NB-r14
// SIZE (1..maxNeighCell-SCPTM-NB-r14)
type ScptmNeighbourcelllistNbR14 struct {
	Value []PciArfcnNbR14 `lb:1,ub:maxNeighCellScptmNbR14`
}
