package ies

import "rrc/utils"

// CQI-ReportPeriodicProcExt-r11-csi-ConfigIndex-r11-setup ::= SEQUENCE
type CqiReportperiodicprocextR11CsiConfigindexR11Setup struct {
	CqiPmiConfigindex2R11 utils.INTEGER  `lb:0,ub:1023`
	RiConfigindex2R11     *utils.INTEGER `lb:0,ub:1023`
}
