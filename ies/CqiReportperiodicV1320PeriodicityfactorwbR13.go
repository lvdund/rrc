package ies

import "rrc/utils"

// CQI-ReportPeriodic-v1320-periodicityFactorWB-r13 ::= ENUMERATED
type CqiReportperiodicV1320PeriodicityfactorwbR13 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportperiodicV1320PeriodicityfactorwbR13EnumeratedNothing = iota
	CqiReportperiodicV1320PeriodicityfactorwbR13EnumeratedN2
	CqiReportperiodicV1320PeriodicityfactorwbR13EnumeratedN4
)
