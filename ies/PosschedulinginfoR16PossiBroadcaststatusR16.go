package ies

import "rrc/utils"

// PosSchedulingInfo-r16-posSI-BroadcastStatus-r16 ::= ENUMERATED
type PosschedulinginfoR16PossiBroadcaststatusR16 struct {
	Value utils.ENUMERATED
}

const (
	PosschedulinginfoR16PossiBroadcaststatusR16EnumeratedNothing = iota
	PosschedulinginfoR16PossiBroadcaststatusR16EnumeratedBroadcasting
	PosschedulinginfoR16PossiBroadcaststatusR16EnumeratedNotbroadcasting
)
