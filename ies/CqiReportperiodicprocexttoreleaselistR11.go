package ies

// CQI-ReportPeriodicProcExtToReleaseList-r11 ::= SEQUENCE OF CQI-ReportPeriodicProcExtId-r11
// SIZE (1..maxCQI-ProcExt-r11)
type CqiReportperiodicprocexttoreleaselistR11 struct {
	Value []CqiReportperiodicprocextidR11 `lb:1,ub:maxCQIProcextR11`
}
