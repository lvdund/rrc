package ies

// MeasObjectToRemoveList ::= SEQUENCE OF MeasObjectId
// SIZE (1..maxNrofObjectId)
type Measobjecttoremovelist struct {
	Value []Measobjectid `lb:1,ub:maxNrofObjectId`
}
