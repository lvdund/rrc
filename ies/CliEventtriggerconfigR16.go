package ies

import "rrc/utils"

// CLI-EventTriggerConfig-r16 ::= SEQUENCE
// Extensible
type CliEventtriggerconfigR16 struct {
	EventidR16        CliEventtriggerconfigR16EventidR16
	ReportintervalR16 Reportinterval
	ReportamountR16   CliEventtriggerconfigR16ReportamountR16
	MaxreportcliR16   utils.INTEGER `lb:0,ub:maxCLIReportR16`
}
