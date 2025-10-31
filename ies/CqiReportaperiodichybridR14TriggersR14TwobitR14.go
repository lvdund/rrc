package ies

import "rrc/utils"

// CQI-ReportAperiodicHybrid-r14-triggers-r14-twoBit-r14 ::= SEQUENCE
type CqiReportaperiodichybridR14TriggersR14TwobitR14 struct {
	Trigger01IndicatorR14 utils.BITSTRING `lb:8,ub:8`
	Trigger10IndicatorR14 utils.BITSTRING `lb:8,ub:8`
	Trigger11IndicatorR14 utils.BITSTRING `lb:8,ub:8`
}
