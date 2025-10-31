package ies

import "rrc/utils"

// EventTriggerConfigInterRAT-eventId-eventB1 ::= SEQUENCE
// Extensible
type EventtriggerconfiginterratEventidEventb1 struct {
	B1Thresholdeutra Meastriggerquantityeutra
	Reportonleave    utils.BOOLEAN
	Hysteresis       Hysteresis
	Timetotrigger    Timetotrigger
}
