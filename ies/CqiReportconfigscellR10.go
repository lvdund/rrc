package ies

import "rrc/utils"

// CQI-ReportConfigSCell-r10 ::= SEQUENCE
type CqiReportconfigscellR10 struct {
	CqiReportmodeaperiodicR10 *CqiReportmodeaperiodic
	NompdschRsEpreOffsetR10   utils.INTEGER `lb:0,ub:6`
	CqiReportperiodicscellR10 *CqiReportperiodicR10
	PmiRiReportR10            *CqiReportconfigscellR10PmiRiReportR10
}
