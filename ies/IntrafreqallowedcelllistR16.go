package ies

// IntraFreqAllowedCellList-r16 ::= SEQUENCE OF PCI-Range
// SIZE (1..maxCellAllowed)
type IntrafreqallowedcelllistR16 struct {
	Value []PciRange `lb:1,ub:maxCellAllowed`
}
