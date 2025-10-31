package ies

import "rrc/utils"

// CLI-PeriodicalReportConfig-r16 ::= SEQUENCE
// Extensible
type CliPeriodicalreportconfigR16 struct {
	ReportintervalR16    Reportinterval
	ReportamountR16      CliPeriodicalreportconfigR16ReportamountR16
	ReportquantitycliR16 MeasreportquantitycliR16
	MaxreportcliR16      utils.INTEGER `lb:0,ub:maxCLIReportR16`
}
