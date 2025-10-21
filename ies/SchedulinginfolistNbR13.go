package ies

import "rrc/utils"

// SchedulingInfoList-NB-r13 ::= SEQUENCE OF SchedulingInfo-NB-r13
// SIZE (1..maxSI-Message-NB-r13)
type SchedulinginfolistNbR13 struct {
	Value utils.Sequence[SchedulinginfoNbR13]
}
