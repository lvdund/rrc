package ies

import "rrc/utils"

// EventTriggerConfig-eventId-eventA3 ::= SEQUENCE
type EventtriggerconfigEventidEventa3 struct {
	A3Offset           Meastriggerquantityoffset
	Reportonleave      utils.BOOLEAN
	Hysteresis         Hysteresis
	Timetotrigger      Timetotrigger
	Useallowedcelllist utils.BOOLEAN
}
