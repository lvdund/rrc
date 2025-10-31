package ies

// MeasResultServFreqListEUTRA-SCG ::= SEQUENCE OF MeasResult2EUTRA
// SIZE (1..maxNrofServingCellsEUTRA)
type MeasresultservfreqlisteutraScg struct {
	Value []Measresult2eutra `lb:1,ub:maxNrofServingCellsEUTRA`
}
