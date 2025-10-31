package ies

import "rrc/utils"

// ReportConfigInterRAT ::= SEQUENCE
// Extensible
type Reportconfiginterrat struct {
	Triggertype    ReportconfiginterratTriggertype
	Maxreportcells utils.INTEGER `lb:0,ub:maxCellReport`
	Reportinterval Reportinterval
	Reportamount   ReportconfiginterratReportamount
}
