package ies

import "rrc/utils"

// ReportConfigEUTRA ::= SEQUENCE
// Extensible
type Reportconfigeutra struct {
	Triggertype     ReportconfigeutraTriggertype
	Triggerquantity ReportconfigeutraTriggerquantity
	Reportquantity  ReportconfigeutraReportquantity
	Maxreportcells  utils.INTEGER `lb:0,ub:maxCellReport`
	Reportinterval  Reportinterval
	Reportamount    ReportconfigeutraReportamount
}
