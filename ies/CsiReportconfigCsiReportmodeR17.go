package ies

import "rrc/utils"

// CSI-ReportConfig-csi-ReportMode-r17 ::= ENUMERATED
type CsiReportconfigCsiReportmodeR17 struct {
	Value utils.ENUMERATED
}

const (
	CsiReportconfigCsiReportmodeR17EnumeratedNothing = iota
	CsiReportconfigCsiReportmodeR17EnumeratedMode1
	CsiReportconfigCsiReportmodeR17EnumeratedMode2
)
