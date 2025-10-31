package ies

import "rrc/utils"

// FailureReportSCG-EUTRA ::= SEQUENCE
// Extensible
type FailurereportscgEutra struct {
	Failuretype              FailurereportscgEutraFailuretype
	Measresultfreqlistmrdc   *Measresultfreqlistfailmrdc
	MeasresultscgFailuremrdc *utils.OCTETSTRING
	LocationinfoR16          *LocationinfoR16 `ext`
}
