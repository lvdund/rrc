package ies

// PCI-RangeIndexList ::= SEQUENCE OF PCI-RangeIndex
// SIZE (1..maxNrofPCI-Ranges)
type PciRangeindexlist struct {
	Value []PciRangeindex `lb:1,ub:maxNrofPCIRanges`
}
