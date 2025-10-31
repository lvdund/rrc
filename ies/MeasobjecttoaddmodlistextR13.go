package ies

// MeasObjectToAddModListExt-r13 ::= SEQUENCE OF MeasObjectToAddModExt-r13
// SIZE (1..maxObjectId)
type MeasobjecttoaddmodlistextR13 struct {
	Value []MeasobjecttoaddmodextR13 `lb:1,ub:maxObjectId`
}
