package ies

import "rrc/utils"

// RSRP-ThresholdsPrachInfoList-r13 ::= SEQUENCE OF RSRP-Range
// SIZE (1..3)
type RsrpThresholdsprachinfolistR13 struct {
	Value []RsrpRange
}
