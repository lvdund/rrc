package ies

import "rrc/utils"

// EventTriggerConfigInterRAT-eventId-eventY1-Relay-r17 ::= SEQUENCE
// Extensible
type EventtriggerconfiginterratEventidEventy1RelayR17 struct {
	Y1Threshold1R17      Meastriggerquantity
	Y1Threshold2RelayR17 SlMeastriggerquantityR16
	ReportonleaveR17     utils.BOOLEAN
	HysteresisR17        Hysteresis
	TimetotriggerR17     Timetotrigger
}
