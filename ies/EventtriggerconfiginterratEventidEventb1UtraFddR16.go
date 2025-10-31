package ies

import "rrc/utils"

// EventTriggerConfigInterRAT-eventId-eventB1-UTRA-FDD-r16 ::= SEQUENCE
// Extensible
type EventtriggerconfiginterratEventidEventb1UtraFddR16 struct {
	B1ThresholdutraFddR16 MeastriggerquantityutraFddR16
	ReportonleaveR16      utils.BOOLEAN
	HysteresisR16         Hysteresis
	TimetotriggerR16      Timetotrigger
}
