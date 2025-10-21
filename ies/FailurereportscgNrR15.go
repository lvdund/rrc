package ies

import "rrc/utils"

// FailureReportSCG-NR-r15 ::= SEQUENCE
// Extensible
type FailurereportscgNrR15 struct {
	FailuretypeR15          utils.ENUMERATED
	MeasresultfreqlistnrR15 *MeasresultfreqlistfailnrR15
	MeasresultscgR15        *utils.OCTETSTRING
}
