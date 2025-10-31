package ies

// UL-ConfigCommonList-NB-v1530 ::= SEQUENCE OF UL-ConfigCommon-NB-v1530
// SIZE (1.. maxNonAnchorCarriers-NB-r14)
type UlConfigcommonlistNbV1530 struct {
	Value []UlConfigcommonNbV1530 `lb:1,ub:maxNonAnchorCarriersNbR14`
}
