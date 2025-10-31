package ies

import "rrc/utils"

// CQI-ReportConfig-v1530-altCQI-Table-1024QAM-r15 ::= ENUMERATED
type CqiReportconfigV1530AltcqiTable1024qamR15 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportconfigV1530AltcqiTable1024qamR15EnumeratedNothing = iota
	CqiReportconfigV1530AltcqiTable1024qamR15EnumeratedAllsubframes
	CqiReportconfigV1530AltcqiTable1024qamR15EnumeratedCsi_Subframeset1
	CqiReportconfigV1530AltcqiTable1024qamR15EnumeratedCsi_Subframeset2
	CqiReportconfigV1530AltcqiTable1024qamR15EnumeratedSpare1
)
