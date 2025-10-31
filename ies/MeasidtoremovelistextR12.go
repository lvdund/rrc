package ies

// MeasIdToRemoveListExt-r12 ::= SEQUENCE OF MeasId-v1250
// SIZE (1..maxMeasId)
type MeasidtoremovelistextR12 struct {
	Value []MeasidV1250 `lb:1,ub:maxMeasId`
}
