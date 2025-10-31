package ies

import "rrc/utils"

// SL-EventTriggerConfig-r16-sl-EventId-r16-eventS1-r16 ::= SEQUENCE
// Extensible
type SlEventtriggerconfigR16SlEventidR16Events1R16 struct {
	S1ThresholdR16     SlMeastriggerquantityR16
	SlReportonleaveR16 utils.BOOLEAN
	SlHysteresisR16    Hysteresis
	SlTimetotriggerR16 Timetotrigger
}
