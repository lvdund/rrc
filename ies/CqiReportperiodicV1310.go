package ies

import "rrc/utils"

// CQI-ReportPeriodic-v1310 ::= SEQUENCE
type CqiReportperiodicV1310 struct {
	CriReportconfigR13                         *CriReportconfigR13
	SimultaneousacknackandcqiFormat4Format5R13 *utils.ENUMERATED
}
