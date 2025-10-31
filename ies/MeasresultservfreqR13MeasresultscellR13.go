package ies

// MeasResultServFreq-r13-measResultSCell-r13 ::= SEQUENCE
type MeasresultservfreqR13MeasresultscellR13 struct {
	RsrpresultscellR13 RsrpRange
	RsrqresultscellR13 RsrqRangeR13
	RsSinrResultR13    *RsSinrRangeR13
}
