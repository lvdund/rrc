package ies

import "rrc/utils"

// CQI-ReportAperiodicProc-r11 ::= SEQUENCE
type CqiReportaperiodicprocR11 struct {
	CqiReportmodeaperiodicR11 CqiReportmodeaperiodic
	Trigger01R11              bool
	Trigger10R11              bool
	Trigger11R11              bool
}
