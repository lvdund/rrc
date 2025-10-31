package ies

import "rrc/utils"

// CQI-ReportAperiodicHybrid-r14-triggers-r14-threeBit-r14 ::= SEQUENCE
type CqiReportaperiodichybridR14TriggersR14ThreebitR14 struct {
	Trigger001IndicatorR14 utils.BITSTRING `lb:32,ub:32`
	Trigger010IndicatorR14 utils.BITSTRING `lb:32,ub:32`
	Trigger011IndicatorR14 utils.BITSTRING `lb:32,ub:32`
	Trigger100IndicatorR14 utils.BITSTRING `lb:32,ub:32`
	Trigger101IndicatorR14 utils.BITSTRING `lb:32,ub:32`
	Trigger110IndicatorR14 utils.BITSTRING `lb:32,ub:32`
	Trigger111IndicatorR14 utils.BITSTRING `lb:32,ub:32`
}
