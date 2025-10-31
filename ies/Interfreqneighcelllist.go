package ies

// InterFreqNeighCellList ::= SEQUENCE OF InterFreqNeighCellInfo
// SIZE (1..maxCellInter)
type Interfreqneighcelllist struct {
	Value []Interfreqneighcellinfo `lb:1,ub:maxCellInter`
}
