package ies

import "rrc/utils"

// CQI-ReportConfig ::= SEQUENCE
type CqiReportconfig struct {
	CqiReportmodeaperiodic *CqiReportmodeaperiodic
	NompdschRsEpreOffset   utils.INTEGER `lb:0,ub:6`
	CqiReportperiodic      *CqiReportperiodic
}
