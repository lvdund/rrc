package ies

// PCI-Range ::= SEQUENCE
type PciRange struct {
	Start Physcellid
	Range *PciRangeRange
}
