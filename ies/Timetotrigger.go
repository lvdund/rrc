package ies

import "rrc/utils"

// TimeToTrigger ::= ENUMERATED
type Timetotrigger struct {
	Value utils.ENUMERATED
}

const (
	TimetotriggerEnumeratedNothing = iota
	TimetotriggerEnumeratedMs0
	TimetotriggerEnumeratedMs40
	TimetotriggerEnumeratedMs64
	TimetotriggerEnumeratedMs80
	TimetotriggerEnumeratedMs100
	TimetotriggerEnumeratedMs128
	TimetotriggerEnumeratedMs160
	TimetotriggerEnumeratedMs256
	TimetotriggerEnumeratedMs320
	TimetotriggerEnumeratedMs480
	TimetotriggerEnumeratedMs512
	TimetotriggerEnumeratedMs640
	TimetotriggerEnumeratedMs1024
	TimetotriggerEnumeratedMs1280
	TimetotriggerEnumeratedMs2560
	TimetotriggerEnumeratedMs5120
)
