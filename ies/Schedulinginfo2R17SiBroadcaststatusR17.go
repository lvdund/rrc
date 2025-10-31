package ies

import "rrc/utils"

// SchedulingInfo2-r17-si-BroadcastStatus-r17 ::= ENUMERATED
type Schedulinginfo2R17SiBroadcaststatusR17 struct {
	Value utils.ENUMERATED
}

const (
	Schedulinginfo2R17SiBroadcaststatusR17EnumeratedNothing = iota
	Schedulinginfo2R17SiBroadcaststatusR17EnumeratedBroadcasting
	Schedulinginfo2R17SiBroadcaststatusR17EnumeratedNotbroadcasting
)
