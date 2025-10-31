package ies

import "rrc/utils"

// CQI-ReportModeAperiodic ::= ENUMERATED
type CqiReportmodeaperiodic struct {
	Value utils.ENUMERATED
}

const (
	CqiReportmodeaperiodicEnumeratedNothing = iota
	CqiReportmodeaperiodicEnumeratedRm12
	CqiReportmodeaperiodicEnumeratedRm20
	CqiReportmodeaperiodicEnumeratedRm22
	CqiReportmodeaperiodicEnumeratedRm30
	CqiReportmodeaperiodicEnumeratedRm31
	CqiReportmodeaperiodicEnumeratedRm32_V1250
	CqiReportmodeaperiodicEnumeratedRm10_V1310
	CqiReportmodeaperiodicEnumeratedRm11_V1310
)
