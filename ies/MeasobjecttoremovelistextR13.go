package ies

// MeasObjectToRemoveListExt-r13 ::= SEQUENCE OF MeasObjectId-v1310
// SIZE (1..maxObjectId)
type MeasobjecttoremovelistextR13 struct {
	Value []MeasobjectidV1310 `lb:1,ub:maxObjectId`
}
