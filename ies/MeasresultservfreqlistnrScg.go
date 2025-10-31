package ies

// MeasResultServFreqListNR-SCG ::= SEQUENCE OF MeasResult2NR
// SIZE (1..maxNrofServingCells)
type MeasresultservfreqlistnrScg struct {
	Value []Measresult2nr `lb:1,ub:maxNrofServingCells`
}
