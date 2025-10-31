package ies

// RedistributionNeighCellList-r13 ::= SEQUENCE OF RedistributionNeighCell-r13
// SIZE (1..maxCellInter)
type RedistributionneighcelllistR13 struct {
	Value []RedistributionneighcellR13 `lb:1,ub:maxCellInter`
}
