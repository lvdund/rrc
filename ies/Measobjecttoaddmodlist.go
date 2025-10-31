package ies

// MeasObjectToAddModList ::= SEQUENCE OF MeasObjectToAddMod
// SIZE (1..maxNrofObjectId)
type Measobjecttoaddmodlist struct {
	Value []Measobjecttoaddmod `lb:1,ub:maxNrofObjectId`
}
