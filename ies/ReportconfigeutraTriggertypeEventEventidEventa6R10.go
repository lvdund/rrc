package ies

import "rrc/utils"

// ReportConfigEUTRA-triggerType-event-eventId-eventA6-r10 ::= SEQUENCE
type ReportconfigeutraTriggertypeEventEventidEventa6R10 struct {
	A6OffsetR10        utils.INTEGER `lb:0,ub:30`
	A6ReportonleaveR10 utils.BOOLEAN
}
