package ies

import "rrc/utils"

// CQI-ReportBothProc-r11-pmi-RI-Report-r11 ::= ENUMERATED
type CqiReportbothprocR11PmiRiReportR11 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportbothprocR11PmiRiReportR11EnumeratedNothing = iota
	CqiReportbothprocR11PmiRiReportR11EnumeratedSetup
)
