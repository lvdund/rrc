package ies

import "rrc/utils"

// CSI-ReportConfig-sharedCMR-r17 ::= ENUMERATED
type CsiReportconfigSharedcmrR17 struct {
	Value utils.ENUMERATED
}

const (
	CsiReportconfigSharedcmrR17EnumeratedNothing = iota
	CsiReportconfigSharedcmrR17EnumeratedEnable
)
