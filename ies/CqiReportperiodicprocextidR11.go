package ies

import "rrc/utils"

// CQI-ReportPeriodicProcExtId-r11 ::= utils.INTEGER (1..maxCQI-ProcExt-r11)
type CqiReportperiodicprocextidR11 struct {
	Value utils.INTEGER `lb:0,ub:maxCQIProcextR11`
}
