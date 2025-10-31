package ies

import "rrc/utils"

// ReportConfigEUTRA-triggerType-event-eventId-eventH1-r15 ::= SEQUENCE
type ReportconfigeutraTriggertypeEventEventidEventh1R15 struct {
	H1ThresholdoffsetR15 utils.INTEGER `lb:0,ub:300`
	H1HysteresisR15      utils.INTEGER `lb:0,ub:16`
}
