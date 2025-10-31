package ies

// InterFreqNeighCellList-v1610 ::= SEQUENCE OF InterFreqNeighCellInfo-v1610
// SIZE (1..maxCellInter)
type InterfreqneighcelllistV1610 struct {
	Value []InterfreqneighcellinfoV1610 `lb:1,ub:maxCellInter`
}
