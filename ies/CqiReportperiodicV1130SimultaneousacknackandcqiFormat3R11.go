package ies

import "rrc/utils"

// CQI-ReportPeriodic-v1130-simultaneousAckNackAndCQI-Format3-r11 ::= ENUMERATED
type CqiReportperiodicV1130SimultaneousacknackandcqiFormat3R11 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportperiodicV1130SimultaneousacknackandcqiFormat3R11EnumeratedNothing = iota
	CqiReportperiodicV1130SimultaneousacknackandcqiFormat3R11EnumeratedSetup
)
