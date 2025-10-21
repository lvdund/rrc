package ies

import "rrc/utils"

// CQI-ReportConfig-v920 ::= SEQUENCE
type CqiReportconfigV920 struct {
	CqiMaskR9     *utils.ENUMERATED
	PmiRiReportR9 *utils.ENUMERATED
}
