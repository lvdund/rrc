package ies

import "rrc/utils"

// ReportConfigEUTRA-triggerType-event-eventId-eventH2-r15 ::= SEQUENCE
type ReportconfigeutraTriggertypeEventEventidEventh2R15 struct {
	H2ThresholdoffsetR15 utils.INTEGER `lb:0,ub:300`
	H2HysteresisR15      utils.INTEGER `lb:0,ub:16`
}
