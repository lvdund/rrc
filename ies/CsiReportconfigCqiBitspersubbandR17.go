package ies

import "rrc/utils"

// CSI-ReportConfig-cqi-BitsPerSubband-r17 ::= ENUMERATED
type CsiReportconfigCqiBitspersubbandR17 struct {
	Value utils.ENUMERATED
}

const (
	CsiReportconfigCqiBitspersubbandR17EnumeratedNothing = iota
	CsiReportconfigCqiBitspersubbandR17EnumeratedBits4
)
