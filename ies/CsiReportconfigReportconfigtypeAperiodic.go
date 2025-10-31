package ies

import "rrc/utils"

// CSI-ReportConfig-reportConfigType-aperiodic ::= SEQUENCE
type CsiReportconfigReportconfigtypeAperiodic struct {
	Reportslotoffsetlist []utils.INTEGER `lb:1,ub:maxNrofULAllocations`
}
