package ies

import "rrc/utils"

// CQI-ReportAperiodicProc-r11 ::= SEQUENCE
type CqiReportaperiodicprocR11 struct {
	CqiReportmodeaperiodicR11 CqiReportmodeaperiodic
	Trigger01R11              utils.BOOLEAN
	Trigger10R11              utils.BOOLEAN
	Trigger11R11              utils.BOOLEAN
}
