package ies

// UL-ConfigCommonListTDD-NB-r15 ::= SEQUENCE OF UL-ConfigCommonTDD-NB-r15
// SIZE (1.. maxNonAnchorCarriers-NB-r14)
type UlConfigcommonlisttddNbR15 struct {
	Value []UlConfigcommontddNbR15 `lb:1,ub:maxNonAnchorCarriersNbR14`
}
