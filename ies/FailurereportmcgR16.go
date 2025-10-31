package ies

import "rrc/utils"

// FailureReportMCG-r16 ::= SEQUENCE
// Extensible
type FailurereportmcgR16 struct {
	FailuretypeR16             *FailurereportmcgR16FailuretypeR16
	MeasresultfreqlisteutraR16 *Measresultlist3eutraR15
	MeasresultfreqlistnrR16    *MeasresultfreqlistfailnrR15
	MeasresultfreqlistgeranR16 *Measresultlist2geranR10
	MeasresultfreqlistutraR16  *Measresultlist2utraR9
	MeasresultscgR16           *utils.OCTETSTRING
}
