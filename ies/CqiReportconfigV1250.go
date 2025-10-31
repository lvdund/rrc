package ies

// CQI-ReportConfig-v1250 ::= SEQUENCE
type CqiReportconfigV1250 struct {
	CsiSubframepatternconfigR12 *CqiReportconfigV1250CsiSubframepatternconfigR12
	CqiReportbothV1250          *CqiReportbothV1250
	CqiReportaperiodicV1250     *CqiReportaperiodicV1250
	AltcqiTableR12              *CqiReportconfigV1250AltcqiTableR12
}
