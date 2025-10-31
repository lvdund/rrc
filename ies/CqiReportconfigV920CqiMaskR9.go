package ies

import "rrc/utils"

// CQI-ReportConfig-v920-cqi-Mask-r9 ::= ENUMERATED
type CqiReportconfigV920CqiMaskR9 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportconfigV920CqiMaskR9EnumeratedNothing = iota
	CqiReportconfigV920CqiMaskR9EnumeratedSetup
)
