package ies

import "rrc/utils"

// CQI-ReportPeriodic-setup ::= SEQUENCE
type CqiReportperiodicSetup struct {
	CqiPucchResourceindex      utils.INTEGER `lb:0,ub:1185`
	CqiPmiConfigindex          utils.INTEGER `lb:0,ub:1023`
	CqiFormatindicatorperiodic CqiReportperiodicSetupCqiFormatindicatorperiodic
	RiConfigindex              *utils.INTEGER `lb:0,ub:1023`
	Simultaneousacknackandcqi  utils.BOOLEAN
}
