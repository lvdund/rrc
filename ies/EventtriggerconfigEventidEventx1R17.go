package ies

import "rrc/utils"

// EventTriggerConfig-eventId-eventX1-r17 ::= SEQUENCE
type EventtriggerconfigEventidEventx1R17 struct {
	X1Threshold1RelayR17  SlMeastriggerquantityR16
	X1Threshold2R17       Meastriggerquantity
	ReportonleaveR17      utils.BOOLEAN
	HysteresisR17         Hysteresis
	TimetotriggerR17      Timetotrigger
	UseallowedcelllistR17 utils.BOOLEAN
}
