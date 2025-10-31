package ies

import "rrc/utils"

// CQI-ReportPeriodicProcExt-r11 ::= SEQUENCE
// Extensible
type CqiReportperiodicprocextR11 struct {
	CqiReportperiodicprocextidR11 CqiReportperiodicprocextidR11
	CqiPmiConfigindexR11          utils.INTEGER `lb:0,ub:1023`
	CqiFormatindicatorperiodicR11 *CqiReportperiodicprocextR11CqiFormatindicatorperiodicR11
	RiConfigindexR11              *utils.INTEGER `lb:0,ub:1023`
	CsiConfigindexR11             *CqiReportperiodicprocextR11CsiConfigindexR11
}
