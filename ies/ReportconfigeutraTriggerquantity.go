package ies

import "rrc/utils"

// ReportConfigEUTRA-triggerQuantity ::= ENUMERATED
type ReportconfigeutraTriggerquantity struct {
	Value utils.ENUMERATED
}

const (
	ReportconfigeutraTriggerquantityEnumeratedNothing = iota
	ReportconfigeutraTriggerquantityEnumeratedRsrp
	ReportconfigeutraTriggerquantityEnumeratedRsrq
)
