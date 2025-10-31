package ies

// MeasIdToRemoveList ::= SEQUENCE OF MeasId
// SIZE (1..maxNrofMeasId)
type Measidtoremovelist struct {
	Value []Measid `lb:1,ub:maxNrofMeasId`
}
