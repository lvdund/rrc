package ies

import "rrc/utils"

// CSI-ReportConfig-dummy ::= ENUMERATED
type CsiReportconfigDummy struct {
	Value utils.ENUMERATED
}

const (
	CsiReportconfigDummyEnumeratedNothing = iota
	CsiReportconfigDummyEnumeratedN1
	CsiReportconfigDummyEnumeratedN2
)
