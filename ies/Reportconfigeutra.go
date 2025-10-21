package ies

import "rrc/utils"

// ReportConfigEUTRA ::= SEQUENCE
// Extensible
type Reportconfigeutra struct {
	Triggertype     interface{}
	Triggerquantity utils.ENUMERATED
	Reportquantity  utils.ENUMERATED
	Maxreportcells  utils.INTEGER
	Reportinterval  Reportinterval
	Reportamount    utils.ENUMERATED
}
