package ies

// IntraFreqNeighCellList-v1610 ::= SEQUENCE OF IntraFreqNeighCellInfo-v1610
// SIZE (1..maxCellIntra)
type IntrafreqneighcelllistV1610 struct {
	Value []IntrafreqneighcellinfoV1610 `lb:1,ub:maxCellIntra`
}
