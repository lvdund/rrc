package ies

// CQI-ReportConfig-v1310 ::= SEQUENCE
type CqiReportconfigV1310 struct {
	CqiReportbothV1310      *CqiReportbothV1310
	CqiReportaperiodicV1310 *CqiReportaperiodicV1310
	CqiReportperiodicV1310  *CqiReportperiodicV1310
}
