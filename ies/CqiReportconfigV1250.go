package ies

import "rrc/utils"

// CQI-ReportConfig-v1250 ::= SEQUENCE
type CqiReportconfigV1250 struct {
	CsiSubframepatternconfigR12 *interface{}
	CqiReportbothV1250          *CqiReportbothV1250
	CqiReportaperiodicV1250     *CqiReportaperiodicV1250
	AltcqiTableR12              *utils.ENUMERATED
}
