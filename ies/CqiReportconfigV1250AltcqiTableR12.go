package ies

import "rrc/utils"

// CQI-ReportConfig-v1250-altCQI-Table-r12 ::= ENUMERATED
type CqiReportconfigV1250AltcqiTableR12 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportconfigV1250AltcqiTableR12EnumeratedNothing = iota
	CqiReportconfigV1250AltcqiTableR12EnumeratedAllsubframes
	CqiReportconfigV1250AltcqiTableR12EnumeratedCsi_Subframeset1
	CqiReportconfigV1250AltcqiTableR12EnumeratedCsi_Subframeset2
	CqiReportconfigV1250AltcqiTableR12EnumeratedSpare1
)
