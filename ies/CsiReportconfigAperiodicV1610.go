package ies

import "rrc/utils"

// CSI-ReportConfig-aperiodic-v1610 ::= SEQUENCE
type CsiReportconfigAperiodicV1610 struct {
	Reportslotoffsetlistdci02R16 *[]utils.INTEGER `lb:1,ub:maxNrofULAllocationsR16`
	Reportslotoffsetlistdci01R16 *[]utils.INTEGER `lb:1,ub:maxNrofULAllocationsR16`
}
