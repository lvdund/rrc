package ies

import "rrc/utils"

// CSI-ReportConfig-aperiodic-v1720 ::= SEQUENCE
type CsiReportconfigAperiodicV1720 struct {
	ReportslotoffsetlistR17      *[]utils.INTEGER `lb:1,ub:maxNrofULAllocationsR16`
	Reportslotoffsetlistdci02R17 *[]utils.INTEGER `lb:1,ub:maxNrofULAllocationsR16`
	Reportslotoffsetlistdci01R17 *[]utils.INTEGER `lb:1,ub:maxNrofULAllocationsR16`
}
