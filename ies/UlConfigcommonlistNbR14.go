package ies

// UL-ConfigCommonList-NB-r14 ::= SEQUENCE OF UL-ConfigCommon-NB-r14
// SIZE (1.. maxNonAnchorCarriers-NB-r14)
type UlConfigcommonlistNbR14 struct {
	Value []UlConfigcommonNbR14 `lb:1,ub:maxNonAnchorCarriersNbR14`
}
