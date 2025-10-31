package ies

import "rrc/utils"

// ReportConfigEUTRA-triggerType-event-eventId-eventA3 ::= SEQUENCE
type ReportconfigeutraTriggertypeEventEventidEventa3 struct {
	A3Offset      utils.INTEGER `lb:0,ub:30`
	Reportonleave utils.BOOLEAN
}
