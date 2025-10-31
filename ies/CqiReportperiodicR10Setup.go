package ies

import "rrc/utils"

// CQI-ReportPeriodic-r10-setup ::= SEQUENCE
type CqiReportperiodicR10Setup struct {
	CqiPucchResourceindexR10      utils.INTEGER  `lb:0,ub:1184`
	CqiPucchResourceindexp1R10    *utils.INTEGER `lb:0,ub:1184`
	CqiPmiConfigindex             utils.INTEGER  `lb:0,ub:1023`
	CqiFormatindicatorperiodicR10 *CqiReportperiodicR10SetupCqiFormatindicatorperiodicR10
	RiConfigindex                 *utils.INTEGER `lb:0,ub:1023`
	Simultaneousacknackandcqi     utils.BOOLEAN
	CqiMaskR9                     *CqiReportperiodicR10SetupCqiMaskR9
	CsiConfigindexR10             *CqiReportperiodicR10SetupCsiConfigindexR10
}
