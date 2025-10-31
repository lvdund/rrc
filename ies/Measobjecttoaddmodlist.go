package ies

// MeasObjectToAddModList ::= SEQUENCE OF MeasObjectToAddMod
// SIZE (1..maxObjectId)
type Measobjecttoaddmodlist struct {
	Value []Measobjecttoaddmod `lb:1,ub:maxObjectId`
}
