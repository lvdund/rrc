package ies

import "rrc/utils"

// ReportConfigId ::= utils.INTEGER (1..maxReportConfigId)
type Reportconfigid struct {
	Value utils.INTEGER `lb:0,ub:maxReportConfigId`
}
