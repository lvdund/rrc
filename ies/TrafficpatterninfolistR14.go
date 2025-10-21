package ies

import "rrc/utils"

// TrafficPatternInfoList-r14 ::= SEQUENCE OF TrafficPatternInfo-r14
// SIZE (1..maxTrafficPattern-r14)
type TrafficpatterninfolistR14 struct {
	Value []TrafficpatterninfoR14
}
