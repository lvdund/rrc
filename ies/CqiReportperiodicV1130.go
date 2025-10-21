package ies

import "rrc/utils"

// CQI-ReportPeriodic-v1130 ::= SEQUENCE
type CqiReportperiodicV1130 struct {
	SimultaneousacknackandcqiFormat3R11      *utils.ENUMERATED
	CqiReportperiodicprocexttoreleaselistR11 *CqiReportperiodicprocexttoreleaselistR11
	CqiReportperiodicprocexttoaddmodlistR11  *CqiReportperiodicprocexttoaddmodlistR11
}
