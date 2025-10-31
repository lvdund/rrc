package ies

// MeasIdToRemoveList ::= SEQUENCE OF MeasId
// SIZE (1..maxMeasId)
type Measidtoremovelist struct {
	Value []Measid `lb:1,ub:maxMeasId`
}
