package ies

// InterFreqNeighCellList-v1710 ::= SEQUENCE OF InterFreqNeighCellInfo-v1710
// SIZE (1..maxCellInter)
type InterfreqneighcelllistV1710 struct {
	Value []InterfreqneighcellinfoV1710 `lb:1,ub:maxCellInter`
}
