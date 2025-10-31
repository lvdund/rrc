package ies

import "rrc/utils"

// EventTriggerConfigInterRAT-eventId-eventB2 ::= SEQUENCE
// Extensible
type EventtriggerconfiginterratEventidEventb2 struct {
	B2Threshold1      Meastriggerquantity
	B2Threshold2eutra Meastriggerquantityeutra
	Reportonleave     utils.BOOLEAN
	Hysteresis        Hysteresis
	Timetotrigger     Timetotrigger
}
