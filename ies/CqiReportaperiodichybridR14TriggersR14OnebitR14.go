package ies

import "rrc/utils"

// CQI-ReportAperiodicHybrid-r14-triggers-r14-oneBit-r14 ::= SEQUENCE
type CqiReportaperiodichybridR14TriggersR14OnebitR14 struct {
	Trigger1IndicatorR14 utils.BITSTRING `lb:8,ub:8`
}
