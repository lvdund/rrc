package ies

// FailureReportSCG-r12 ::= SEQUENCE
// Extensible
type FailurereportscgR12 struct {
	FailuretypeR12            FailurereportscgR12FailuretypeR12
	MeasresultservfreqlistR12 *MeasresultservfreqlistR10
	MeasresultneighcellsR12   *Measresultlist2eutraR9
}
