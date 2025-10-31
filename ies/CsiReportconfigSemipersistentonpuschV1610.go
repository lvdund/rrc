package ies

import "rrc/utils"

// CSI-ReportConfig-semiPersistentOnPUSCH-v1610 ::= SEQUENCE
type CsiReportconfigSemipersistentonpuschV1610 struct {
	Reportslotoffsetlistdci02R16 *[]utils.INTEGER `lb:1,ub:maxNrofULAllocationsR16`
	Reportslotoffsetlistdci01R16 *[]utils.INTEGER `lb:1,ub:maxNrofULAllocationsR16`
}
