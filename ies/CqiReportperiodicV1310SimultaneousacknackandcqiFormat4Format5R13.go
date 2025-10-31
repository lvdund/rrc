package ies

import "rrc/utils"

// CQI-ReportPeriodic-v1310-simultaneousAckNackAndCQI-Format4-Format5-r13 ::= ENUMERATED
type CqiReportperiodicV1310SimultaneousacknackandcqiFormat4Format5R13 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportperiodicV1310SimultaneousacknackandcqiFormat4Format5R13EnumeratedNothing = iota
	CqiReportperiodicV1310SimultaneousacknackandcqiFormat4Format5R13EnumeratedSetup
)
