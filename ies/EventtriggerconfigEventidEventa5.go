package ies

import "rrc/utils"

// EventTriggerConfig-eventId-eventA5 ::= SEQUENCE
type EventtriggerconfigEventidEventa5 struct {
	A5Threshold1       Meastriggerquantity
	A5Threshold2       Meastriggerquantity
	Reportonleave      utils.BOOLEAN
	Hysteresis         Hysteresis
	Timetotrigger      Timetotrigger
	Useallowedcelllist utils.BOOLEAN
}
