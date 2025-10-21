package ies

import "rrc/utils"

// CQI-ReportModeAperiodic ::= ENUMERATED
type CqiReportmodeaperiodic struct {
	Value utils.ENUMERATED
}

const (
	CqiReportmodeaperiodicRm12      = 0
	CqiReportmodeaperiodicRm20      = 1
	CqiReportmodeaperiodicRm22      = 2
	CqiReportmodeaperiodicRm30      = 3
	CqiReportmodeaperiodicRm31      = 4
	CqiReportmodeaperiodicRm32V1250 = 5
	CqiReportmodeaperiodicRm10V1310 = 6
	CqiReportmodeaperiodicRm11V1310 = 7
)
