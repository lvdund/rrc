package ies

import "rrc/utils"

// CQI-ReportPeriodicProcExtToAddModList-r11 ::= SEQUENCE OF CQI-ReportPeriodicProcExt-r11
// SIZE (1..maxCQI-ProcExt-r11)
type CqiReportperiodicprocexttoaddmodlistR11 struct {
	Value utils.Sequence[CqiReportperiodicprocextR11]
}
