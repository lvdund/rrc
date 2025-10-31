package ies

import "rrc/utils"

// ReportConfigInterRAT-triggerType-event-eventId-eventB1-NR-r15 ::= SEQUENCE
type ReportconfiginterratTriggertypeEventEventidEventb1NrR15 struct {
	B1ThresholdnrR15 ThresholdnrR15
	ReportonleaveR15 utils.BOOLEAN
}
