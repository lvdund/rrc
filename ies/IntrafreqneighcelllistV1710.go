package ies

// IntraFreqNeighCellList-v1710 ::= SEQUENCE OF IntraFreqNeighCellInfo-v1710
// SIZE (1..maxCellIntra)
type IntrafreqneighcelllistV1710 struct {
	Value []IntrafreqneighcellinfoV1710 `lb:1,ub:maxCellIntra`
}
