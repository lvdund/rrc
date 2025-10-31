package ies

import "rrc/utils"

// CQI-ReportConfig-r10-pmi-RI-Report-r9 ::= ENUMERATED
type CqiReportconfigR10PmiRiReportR9 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportconfigR10PmiRiReportR9EnumeratedNothing = iota
	CqiReportconfigR10PmiRiReportR9EnumeratedSetup
)
