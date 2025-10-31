package ies

import "rrc/utils"

// EventTriggerConfig-eventId-eventA1 ::= SEQUENCE
type EventtriggerconfigEventidEventa1 struct {
	A1Threshold   Meastriggerquantity
	Reportonleave utils.BOOLEAN
	Hysteresis    Hysteresis
	Timetotrigger Timetotrigger
}
