package ies

import "rrc/utils"

// CSI-ReportConfig-numberOfSingleTRP-CSI-Mode1-r17 ::= ENUMERATED
type CsiReportconfigNumberofsingletrpCsiMode1R17 struct {
	Value utils.ENUMERATED
}

const (
	CsiReportconfigNumberofsingletrpCsiMode1R17EnumeratedNothing = iota
	CsiReportconfigNumberofsingletrpCsiMode1R17EnumeratedN0
	CsiReportconfigNumberofsingletrpCsiMode1R17EnumeratedN1
	CsiReportconfigNumberofsingletrpCsiMode1R17EnumeratedN2
)
