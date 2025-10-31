package ies

// CQI-ReportBoth-r11 ::= SEQUENCE
type CqiReportbothR11 struct {
	CsiImConfigtoreleaselistR11 *CsiImConfigtoreleaselistR11
	CsiImConfigtoaddmodlistR11  *CsiImConfigtoaddmodlistR11
	CsiProcesstoreleaselistR11  *CsiProcesstoreleaselistR11
	CsiProcesstoaddmodlistR11   *CsiProcesstoaddmodlistR11
}
