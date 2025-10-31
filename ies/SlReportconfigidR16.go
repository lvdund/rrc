package ies

import "rrc/utils"

// SL-ReportConfigId-r16 ::= utils.INTEGER (1..maxNrofSL-ReportConfigId-r16)
type SlReportconfigidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSLReportconfigidR16`
}
