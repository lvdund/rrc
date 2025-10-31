package ies

import "rrc/utils"

// CQI-ReportAperiodic-v1250-setup-aperiodicCSI-Trigger-v1250 ::= SEQUENCE
type CqiReportaperiodicV1250SetupAperiodiccsiTriggerV1250 struct {
	TriggerSubframesetindicatorR12  CqiReportaperiodicV1250SetupAperiodiccsiTriggerV1250TriggerSubframesetindicatorR12
	Trigger1SubframesetindicatorR12 utils.BITSTRING `lb:8,ub:8`
	Trigger2SubframesetindicatorR12 utils.BITSTRING `lb:8,ub:8`
}
