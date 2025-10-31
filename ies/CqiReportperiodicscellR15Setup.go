package ies

import "rrc/utils"

// CQI-ReportPeriodicSCell-r15-setup ::= SEQUENCE
type CqiReportperiodicscellR15Setup struct {
	CqiPmiConfigindexdormantR15  utils.INTEGER  `lb:0,ub:1023`
	RiConfigindexdormantR15      *utils.INTEGER `lb:0,ub:1023`
	CsiSubframepatterndormantR15 *CqiReportperiodicscellR15SetupCsiSubframepatterndormantR15
	CqiFormatindicatordormantR15 *CqiReportperiodicscellR15SetupCqiFormatindicatordormantR15
}
