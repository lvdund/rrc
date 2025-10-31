package ies

// PLMN-RAN-AreaCellList ::= SEQUENCE OF PLMN-RAN-AreaCell
// SIZE (1.. maxPLMNIdentities)
type PlmnRanAreacelllist struct {
	Value []PlmnRanAreacell `lb:1,ub:maxPLMNIdentities`
}
