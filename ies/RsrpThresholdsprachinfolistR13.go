package ies

// RSRP-ThresholdsPrachInfoList-r13 ::= SEQUENCE OF RSRP-Range
// SIZE (1..3)
type RsrpThresholdsprachinfolistR13 struct {
	Value []RsrpRange `lb:1,ub:3`
}
