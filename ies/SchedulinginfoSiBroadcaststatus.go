package ies

import "rrc/utils"

// SchedulingInfo-si-BroadcastStatus ::= ENUMERATED
type SchedulinginfoSiBroadcaststatus struct {
	Value utils.ENUMERATED
}

const (
	SchedulinginfoSiBroadcaststatusEnumeratedNothing = iota
	SchedulinginfoSiBroadcaststatusEnumeratedBroadcasting
	SchedulinginfoSiBroadcaststatusEnumeratedNotbroadcasting
)
