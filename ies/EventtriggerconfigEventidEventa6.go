package ies

import "rrc/utils"

// EventTriggerConfig-eventId-eventA6 ::= SEQUENCE
type EventtriggerconfigEventidEventa6 struct {
	A6Offset           Meastriggerquantityoffset
	Reportonleave      utils.BOOLEAN
	Hysteresis         Hysteresis
	Timetotrigger      Timetotrigger
	Useallowedcelllist utils.BOOLEAN
}
