package ies

import "rrc/utils"

// CSI-ReportConfig-reportFreqConfiguration-pmi-FormatIndicator ::= ENUMERATED
type CsiReportconfigReportfreqconfigurationPmiFormatindicator struct {
	Value utils.ENUMERATED
}

const (
	CsiReportconfigReportfreqconfigurationPmiFormatindicatorEnumeratedNothing = iota
	CsiReportconfigReportfreqconfigurationPmiFormatindicatorEnumeratedWidebandpmi
	CsiReportconfigReportfreqconfigurationPmiFormatindicatorEnumeratedSubbandpmi
)
