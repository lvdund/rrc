package ies

// IntraFreqNeighCellList ::= SEQUENCE OF IntraFreqNeighCellInfo
// SIZE (1..maxCellIntra)
type Intrafreqneighcelllist struct {
	Value []Intrafreqneighcellinfo `lb:1,ub:maxCellIntra`
}
