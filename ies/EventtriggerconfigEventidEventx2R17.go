package ies

import "rrc/utils"

// EventTriggerConfig-eventId-eventX2-r17 ::= SEQUENCE
type EventtriggerconfigEventidEventx2R17 struct {
	X2ThresholdRelayR17 SlMeastriggerquantityR16
	ReportonleaveR17    utils.BOOLEAN
	HysteresisR17       Hysteresis
	TimetotriggerR17    Timetotrigger
}
