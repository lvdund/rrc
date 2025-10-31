package ies

// IntraFreqExcludedCellList ::= SEQUENCE OF PCI-Range
// SIZE (1..maxCellExcluded)
type Intrafreqexcludedcelllist struct {
	Value []PciRange `lb:1,ub:maxCellExcluded`
}
