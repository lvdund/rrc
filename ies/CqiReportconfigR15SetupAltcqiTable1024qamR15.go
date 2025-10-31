package ies

import "rrc/utils"

// CQI-ReportConfig-r15-setup-altCQI-Table-1024QAM-r15 ::= ENUMERATED
type CqiReportconfigR15SetupAltcqiTable1024qamR15 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportconfigR15SetupAltcqiTable1024qamR15EnumeratedNothing = iota
	CqiReportconfigR15SetupAltcqiTable1024qamR15EnumeratedAllsubframes
	CqiReportconfigR15SetupAltcqiTable1024qamR15EnumeratedCsi_Subframeset1
	CqiReportconfigR15SetupAltcqiTable1024qamR15EnumeratedCsi_Subframeset2
	CqiReportconfigR15SetupAltcqiTable1024qamR15EnumeratedSpare1
)
