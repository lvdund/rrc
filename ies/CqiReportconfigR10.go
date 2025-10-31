package ies

import "rrc/utils"

// CQI-ReportConfig-r10 ::= SEQUENCE
type CqiReportconfigR10 struct {
	CqiReportaperiodicR10       *CqiReportaperiodicR10
	NompdschRsEpreOffset        utils.INTEGER `lb:0,ub:6`
	CqiReportperiodicR10        *CqiReportperiodicR10
	PmiRiReportR9               *CqiReportconfigR10PmiRiReportR9
	CsiSubframepatternconfigR10 *CqiReportconfigR10CsiSubframepatternconfigR10
}
