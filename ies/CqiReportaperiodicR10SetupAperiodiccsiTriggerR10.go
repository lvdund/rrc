package ies

import "rrc/utils"

// CQI-ReportAperiodic-r10-setup-aperiodicCSI-Trigger-r10 ::= SEQUENCE
type CqiReportaperiodicR10SetupAperiodiccsiTriggerR10 struct {
	Trigger1R10 utils.BITSTRING `lb:8,ub:8`
	Trigger2R10 utils.BITSTRING `lb:8,ub:8`
}
