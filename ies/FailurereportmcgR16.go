package ies

import "rrc/utils"

// FailureReportMCG-r16 ::= SEQUENCE
// Extensible
type FailurereportmcgR16 struct {
	FailuretypeR16               *FailurereportmcgR16FailuretypeR16
	MeasresultfreqlistR16        *Measresultlist2nr
	MeasresultfreqlisteutraR16   *Measresultlist2eutra
	MeasresultscgR16             *utils.OCTETSTRING
	MeasresultscgEutraR16        *utils.OCTETSTRING
	MeasresultfreqlistutraFddR16 *Measresultlist2utra
}
