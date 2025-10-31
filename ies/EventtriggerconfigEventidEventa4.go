package ies

import "rrc/utils"

// EventTriggerConfig-eventId-eventA4 ::= SEQUENCE
type EventtriggerconfigEventidEventa4 struct {
	A4Threshold        Meastriggerquantity
	Reportonleave      utils.BOOLEAN
	Hysteresis         Hysteresis
	Timetotrigger      Timetotrigger
	Useallowedcelllist utils.BOOLEAN
}
