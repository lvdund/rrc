package ies

import "rrc/utils"

// EventTriggerConfig-eventId-eventA2 ::= SEQUENCE
type EventtriggerconfigEventidEventa2 struct {
	A2Threshold   Meastriggerquantity
	Reportonleave utils.BOOLEAN
	Hysteresis    Hysteresis
	Timetotrigger Timetotrigger
}
