package ies

import "rrc/utils"

// CQI-ReportPeriodic-r10-setup-csi-ConfigIndex-r10-setup ::= SEQUENCE
type CqiReportperiodicR10SetupCsiConfigindexR10Setup struct {
	CqiPmiConfigindex2R10 utils.INTEGER  `lb:0,ub:1023`
	RiConfigindex2R10     *utils.INTEGER `lb:0,ub:1023`
}
