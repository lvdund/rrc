package ies

import "rrc/utils"

// EventTriggerConfigInterRAT-eventId-eventY2-Relay-r17 ::= SEQUENCE
// Extensible
type EventtriggerconfiginterratEventidEventy2RelayR17 struct {
	Y2ThresholdRelayR17 SlMeastriggerquantityR16
	ReportonleaveR17    utils.BOOLEAN
	HysteresisR17       Hysteresis
	TimetotriggerR17    Timetotrigger
}
