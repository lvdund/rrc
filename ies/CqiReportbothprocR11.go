package ies

// CQI-ReportBothProc-r11 ::= SEQUENCE
type CqiReportbothprocR11 struct {
	RiRefCsiProcessidR11 *CsiProcessidR11
	PmiRiReportR11       *CqiReportbothprocR11PmiRiReportR11
}
