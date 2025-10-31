package ies

import "rrc/utils"

// CSI-ReportConfigId ::= utils.INTEGER (0..maxNrofCSI-ReportConfigurations-1)
type CsiReportconfigid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofCSIReportconfigurations1`
}
