package ies

import "rrc/utils"

// CLI-EventTriggerConfig-r16-eventId-r16-eventI1-r16 ::= SEQUENCE
type CliEventtriggerconfigR16EventidR16Eventi1R16 struct {
	I1ThresholdR16   MeastriggerquantitycliR16
	ReportonleaveR16 utils.BOOLEAN
	HysteresisR16    Hysteresis
	TimetotriggerR16 Timetotrigger
}
