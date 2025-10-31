package ies

import "rrc/utils"

// CSI-ReportConfig-subbandSize ::= ENUMERATED
type CsiReportconfigSubbandsize struct {
	Value utils.ENUMERATED
}

const (
	CsiReportconfigSubbandsizeEnumeratedNothing = iota
	CsiReportconfigSubbandsizeEnumeratedValue1
	CsiReportconfigSubbandsizeEnumeratedValue2
)
