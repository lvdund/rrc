package ies

import "rrc/utils"

// FailureReportSCG-r12 ::= SEQUENCE
// Extensible
type FailurereportscgR12 struct {
	FailuretypeR12            utils.ENUMERATED
	MeasresultservfreqlistR12 *MeasresultservfreqlistR10
	MeasresultneighcellsR12   *Measresultlist2eutraR9
}
