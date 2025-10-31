package ies

// MeasResultServFreq-r13-measResultBestNeighCell-r13 ::= SEQUENCE
type MeasresultservfreqR13MeasresultbestneighcellR13 struct {
	PhyscellidR13      Physcellid
	RsrpresultncellR13 RsrpRange
	RsrqresultncellR13 RsrqRangeR13
	RsSinrResultR13    *RsSinrRangeR13
}
