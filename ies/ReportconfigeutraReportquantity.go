package ies

import "rrc/utils"

// ReportConfigEUTRA-reportQuantity ::= ENUMERATED
type ReportconfigeutraReportquantity struct {
	Value utils.ENUMERATED
}

const (
	ReportconfigeutraReportquantityEnumeratedNothing = iota
	ReportconfigeutraReportquantityEnumeratedSameastriggerquantity
	ReportconfigeutraReportquantityEnumeratedBoth
)
