package ies

import "rrc/utils"

// ReportConfigEUTRA-triggerType-event-eventId-eventC1-r12 ::= SEQUENCE
type ReportconfigeutraTriggertypeEventEventidEventc1R12 struct {
	C1ThresholdR12     ThresholdeutraV1250
	C1ReportonleaveR12 utils.BOOLEAN
}
