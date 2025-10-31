package ies

// EUTRA-FreqExcludedCellList ::= SEQUENCE OF EUTRA-PhysCellIdRange
// SIZE (1..maxEUTRA-CellExcluded)
type EutraFreqexcludedcelllist struct {
	Value []EutraPhyscellidrange `lb:1,ub:maxEUTRACellexcluded`
}
