package ies

// MeasResultServMOList ::= SEQUENCE OF MeasResultServMO
// SIZE (1..maxNrofServingCells)
type Measresultservmolist struct {
	Value []Measresultservmo `lb:1,ub:maxNrofServingCells`
}
