package ies

// InterFreqExcludedCellList ::= SEQUENCE OF PCI-Range
// SIZE (1..maxCellExcluded)
type Interfreqexcludedcelllist struct {
	Value []PciRange `lb:1,ub:maxCellExcluded`
}
