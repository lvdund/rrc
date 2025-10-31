package ies

import "rrc/utils"

// CQI-ReportConfig-v920-pmi-RI-Report-r9 ::= ENUMERATED
type CqiReportconfigV920PmiRiReportR9 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportconfigV920PmiRiReportR9EnumeratedNothing = iota
	CqiReportconfigV920PmiRiReportR9EnumeratedSetup
)
