package ies

import "rrc/utils"

// ReportConfigEUTRA-triggerType-event-eventId-eventC2-r12 ::= SEQUENCE
type ReportconfigeutraTriggertypeEventEventidEventc2R12 struct {
	C2RefcsiRsR12      MeascsiRsIdR12
	C2OffsetR12        utils.INTEGER `lb:0,ub:30`
	C2ReportonleaveR12 utils.BOOLEAN
}
