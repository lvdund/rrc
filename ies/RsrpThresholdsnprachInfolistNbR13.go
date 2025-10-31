package ies

// RSRP-ThresholdsNPRACH-InfoList-NB-r13 ::= SEQUENCE OF RSRP-Range
// SIZE (1..2)
type RsrpThresholdsnprachInfolistNbR13 struct {
	Value []RsrpRange `lb:1,ub:2`
}
