package ies

import "rrc/utils"

// ReportConfigInterRAT-triggerType-event-eventId-eventB2-NR-r15 ::= SEQUENCE
type ReportconfiginterratTriggertypeEventEventidEventb2NrR15 struct {
	B2Threshold1R15   Thresholdeutra
	B2Threshold2nrR15 ThresholdnrR15
	ReportonleaveR15  utils.BOOLEAN
}
