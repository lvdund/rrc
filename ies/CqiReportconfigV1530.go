package ies

import "rrc/utils"

// CQI-ReportConfig-v1530 ::= SEQUENCE
type CqiReportconfigV1530 struct {
	AltcqiTable1024qamR15 *utils.ENUMERATED
}
