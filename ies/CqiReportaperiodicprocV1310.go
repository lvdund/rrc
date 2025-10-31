package ies

import "rrc/utils"

// CQI-ReportAperiodicProc-v1310 ::= SEQUENCE
type CqiReportaperiodicprocV1310 struct {
	Trigger001R13 utils.BOOLEAN
	Trigger010R13 utils.BOOLEAN
	Trigger011R13 utils.BOOLEAN
	Trigger100R13 utils.BOOLEAN
	Trigger101R13 utils.BOOLEAN
	Trigger110R13 utils.BOOLEAN
	Trigger111R13 utils.BOOLEAN
}
