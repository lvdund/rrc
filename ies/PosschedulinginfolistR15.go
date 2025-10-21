package ies

import "rrc/utils"

// PosSchedulingInfoList-r15 ::= SEQUENCE OF PosSchedulingInfo-r15
// SIZE (1..maxSI-Message)
type PosschedulinginfolistR15 struct {
	Value utils.Sequence[PosschedulinginfoR15]
}
