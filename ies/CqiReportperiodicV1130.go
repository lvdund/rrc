package ies

// CQI-ReportPeriodic-v1130 ::= SEQUENCE
type CqiReportperiodicV1130 struct {
	SimultaneousacknackandcqiFormat3R11      *CqiReportperiodicV1130SimultaneousacknackandcqiFormat3R11
	CqiReportperiodicprocexttoreleaselistR11 *CqiReportperiodicprocexttoreleaselistR11
	CqiReportperiodicprocexttoaddmodlistR11  *CqiReportperiodicprocexttoaddmodlistR11
}
