package ies

import "rrc/utils"

// CQI-ReportConfig-r10 ::= SEQUENCE
type CqiReportconfigR10 struct {
	CqiReportaperiodicR10       *CqiReportaperiodicR10
	NompdschRsEpreOffset        utils.INTEGER
	CqiReportperiodicR10        *CqiReportperiodicR10
	PmiRiReportR9               *utils.ENUMERATED
	CsiSubframepatternconfigR10 *interface{}
}
