package ies

import "rrc/utils"

// TrafficPatternInfoList-v1530 ::= SEQUENCE OF TrafficPatternInfo-v1530
// SIZE (1..maxTrafficPattern-r14)
type TrafficpatterninfolistV1530 struct {
	Value []TrafficpatterninfoV1530
}
