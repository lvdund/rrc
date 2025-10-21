package ies

import "rrc/utils"

// CQI-ReportConfig ::= SEQUENCE
type CqiReportconfig struct {
	CqiReportmodeaperiodic *CqiReportmodeaperiodic
	NompdschRsEpreOffset   utils.INTEGER
	CqiReportperiodic      *CqiReportperiodic
}
