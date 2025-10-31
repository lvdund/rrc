package ies

import "rrc/utils"

// CSI-ReportConfig-reportFreqConfiguration-cqi-FormatIndicator ::= ENUMERATED
type CsiReportconfigReportfreqconfigurationCqiFormatindicator struct {
	Value utils.ENUMERATED
}

const (
	CsiReportconfigReportfreqconfigurationCqiFormatindicatorEnumeratedNothing = iota
	CsiReportconfigReportfreqconfigurationCqiFormatindicatorEnumeratedWidebandcqi
	CsiReportconfigReportfreqconfigurationCqiFormatindicatorEnumeratedSubbandcqi
)
