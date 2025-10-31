package ies

// EUTRA-FreqNeighCellList ::= SEQUENCE OF EUTRA-FreqNeighCellInfo
// SIZE (1..maxCellEUTRA)
type EutraFreqneighcelllist struct {
	Value []EutraFreqneighcellinfo `lb:1,ub:maxCellEUTRA`
}
