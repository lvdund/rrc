package ies

// MeasResultSCG-Failure ::= SEQUENCE
// Extensible
type MeasresultscgFailure struct {
	Measresultpermolist Measresultlist2nr
	LocationinfoR16     *LocationinfoR16 `ext`
}
