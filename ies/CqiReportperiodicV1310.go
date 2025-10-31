package ies

// CQI-ReportPeriodic-v1310 ::= SEQUENCE
type CqiReportperiodicV1310 struct {
	CriReportconfigR13                         *CriReportconfigR13
	SimultaneousacknackandcqiFormat4Format5R13 *CqiReportperiodicV1310SimultaneousacknackandcqiFormat4Format5R13
}
