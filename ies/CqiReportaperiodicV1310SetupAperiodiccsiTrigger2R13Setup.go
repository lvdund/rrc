package ies

import "rrc/utils"

// CQI-ReportAperiodic-v1310-setup-aperiodicCSI-Trigger2-r13-setup ::= SEQUENCE
type CqiReportaperiodicV1310SetupAperiodiccsiTrigger2R13Setup struct {
	Trigger1SubframesetindicatorR13 utils.BITSTRING `lb:32,ub:32`
	Trigger2SubframesetindicatorR13 utils.BITSTRING `lb:32,ub:32`
	Trigger3SubframesetindicatorR13 utils.BITSTRING `lb:32,ub:32`
	Trigger4SubframesetindicatorR13 utils.BITSTRING `lb:32,ub:32`
	Trigger5SubframesetindicatorR13 utils.BITSTRING `lb:32,ub:32`
	Trigger6SubframesetindicatorR13 utils.BITSTRING `lb:32,ub:32`
}
