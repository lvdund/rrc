package ies

// SCPTM-NeighbourCellList-r13 ::= SEQUENCE OF PCI-ARFCN-r13
// SIZE (1..maxNeighCell-SCPTM-r13)
type ScptmNeighbourcelllistR13 struct {
	Value []PciArfcnR13 `lb:1,ub:maxNeighCellScptmR13`
}
