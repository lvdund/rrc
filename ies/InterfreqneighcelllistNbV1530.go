package ies

// InterFreqNeighCellList-NB-v1530 ::= SEQUENCE OF InterFreqNeighCellInfo-NB-v1530
// SIZE (1..maxCellInter)
type InterfreqneighcelllistNbV1530 struct {
	Value []InterfreqneighcellinfoNbV1530 `lb:1,ub:maxCellInter`
}
