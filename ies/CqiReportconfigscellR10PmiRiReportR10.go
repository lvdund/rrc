package ies

import "rrc/utils"

// CQI-ReportConfigSCell-r10-pmi-RI-Report-r10 ::= ENUMERATED
type CqiReportconfigscellR10PmiRiReportR10 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportconfigscellR10PmiRiReportR10EnumeratedNothing = iota
	CqiReportconfigscellR10PmiRiReportR10EnumeratedSetup
)
