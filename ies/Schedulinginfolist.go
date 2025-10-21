package ies

import "rrc/utils"

// SchedulingInfoList ::= SEQUENCE OF SchedulingInfo
// SIZE (1..maxSI-Message)
type Schedulinginfolist struct {
	Value utils.Sequence[Schedulinginfo]
}
