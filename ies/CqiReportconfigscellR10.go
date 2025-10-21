package ies

import "rrc/utils"

// CQI-ReportConfigSCell-r10 ::= SEQUENCE
type CqiReportconfigscellR10 struct {
	CqiReportmodeaperiodicR10 *CqiReportmodeaperiodic
	NompdschRsEpreOffsetR10   utils.INTEGER
	CqiReportperiodicscellR10 *CqiReportperiodicR10
	PmiRiReportR10            *utils.ENUMERATED
}
