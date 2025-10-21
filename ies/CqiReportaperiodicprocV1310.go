package ies

import "rrc/utils"

// CQI-ReportAperiodicProc-v1310 ::= SEQUENCE
type CqiReportaperiodicprocV1310 struct {
	Trigger001R13 bool
	Trigger010R13 bool
	Trigger011R13 bool
	Trigger100R13 bool
	Trigger101R13 bool
	Trigger110R13 bool
	Trigger111R13 bool
}
