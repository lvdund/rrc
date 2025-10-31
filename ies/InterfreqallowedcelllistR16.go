package ies

// InterFreqAllowedCellList-r16 ::= SEQUENCE OF PCI-Range
// SIZE (1..maxCellAllowed)
type InterfreqallowedcelllistR16 struct {
	Value []PciRange `lb:1,ub:maxCellAllowed`
}
