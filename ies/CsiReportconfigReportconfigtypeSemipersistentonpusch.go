package ies

import "rrc/utils"

// CSI-ReportConfig-reportConfigType-semiPersistentOnPUSCH ::= SEQUENCE
type CsiReportconfigReportconfigtypeSemipersistentonpusch struct {
	Reportslotconfig     CsiReportconfigReportconfigtypeSemipersistentonpuschReportslotconfig
	Reportslotoffsetlist []utils.INTEGER `lb:1,ub:maxNrofULAllocations`
	P0alpha              P0PuschAlphasetid
}
