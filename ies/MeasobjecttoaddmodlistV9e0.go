package ies

// MeasObjectToAddModList-v9e0 ::= SEQUENCE OF MeasObjectToAddMod-v9e0
// SIZE (1..maxObjectId)
type MeasobjecttoaddmodlistV9e0 struct {
	Value []MeasobjecttoaddmodV9e0 `lb:1,ub:maxObjectId`
}
