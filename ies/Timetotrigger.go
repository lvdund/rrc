package ies

import "rrc/utils"

// TimeToTrigger ::= ENUMERATED
type Timetotrigger struct {
	Value utils.ENUMERATED
}

const (
	TimetotriggerMs0    = 0
	TimetotriggerMs40   = 1
	TimetotriggerMs64   = 2
	TimetotriggerMs80   = 3
	TimetotriggerMs100  = 4
	TimetotriggerMs128  = 5
	TimetotriggerMs160  = 6
	TimetotriggerMs256  = 7
	TimetotriggerMs320  = 8
	TimetotriggerMs480  = 9
	TimetotriggerMs512  = 10
	TimetotriggerMs640  = 11
	TimetotriggerMs1024 = 12
	TimetotriggerMs1280 = 13
	TimetotriggerMs2560 = 14
	TimetotriggerMs5120 = 15
)
