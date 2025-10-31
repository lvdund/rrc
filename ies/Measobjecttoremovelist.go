package ies

// MeasObjectToRemoveList ::= SEQUENCE OF MeasObjectId
// SIZE (1..maxObjectId)
type Measobjecttoremovelist struct {
	Value []Measobjectid `lb:1,ub:maxObjectId`
}
