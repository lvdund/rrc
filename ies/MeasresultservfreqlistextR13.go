package ies

// MeasResultServFreqListExt-r13 ::= SEQUENCE OF MeasResultServFreq-r13
// SIZE (1..maxServCell-r13)
type MeasresultservfreqlistextR13 struct {
	Value []MeasresultservfreqR13 `lb:1,ub:maxServCellR13`
}
