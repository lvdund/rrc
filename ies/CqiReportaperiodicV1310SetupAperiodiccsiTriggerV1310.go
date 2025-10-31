package ies

import "rrc/utils"

// CQI-ReportAperiodic-v1310-setup-aperiodicCSI-Trigger-v1310 ::= SEQUENCE
type CqiReportaperiodicV1310SetupAperiodiccsiTriggerV1310 struct {
	Trigger1R13 utils.BITSTRING `lb:32,ub:32`
	Trigger2R13 utils.BITSTRING `lb:32,ub:32`
	Trigger3R13 utils.BITSTRING `lb:32,ub:32`
	Trigger4R13 utils.BITSTRING `lb:32,ub:32`
	Trigger5R13 utils.BITSTRING `lb:32,ub:32`
	Trigger6R13 utils.BITSTRING `lb:32,ub:32`
}
