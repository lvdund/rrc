package ies

// PLMN-RAN-AreaCellList-r15 ::= SEQUENCE OF PLMN-RAN-AreaCell-r15
// SIZE (1..maxPLMN-r15)
type PlmnRanAreacelllistR15 struct {
	Value []PlmnRanAreacellR15 `lb:1,ub:maxPLMNR15`
}
