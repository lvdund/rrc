package ies

import "rrc/utils"

// EventTriggerConfigInterRAT-eventId-eventB2-UTRA-FDD-r16 ::= SEQUENCE
// Extensible
type EventtriggerconfiginterratEventidEventb2UtraFddR16 struct {
	B2Threshold1R16        Meastriggerquantity
	B2Threshold2utraFddR16 MeastriggerquantityutraFddR16
	ReportonleaveR16       utils.BOOLEAN
	HysteresisR16          Hysteresis
	TimetotriggerR16       Timetotrigger
}
