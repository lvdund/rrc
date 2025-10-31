package ies

// CondReconfigExecCondSCG-r17 ::= SEQUENCE OF MeasId
// SIZE (1..2)
type CondreconfigexeccondscgR17 struct {
	Value []Measid `lb:1,ub:2`
}
