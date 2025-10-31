package ies

// MeasIdToAddModListExt-r12 ::= SEQUENCE OF MeasIdToAddModExt-r12
// SIZE (1..maxMeasId)
type MeasidtoaddmodlistextR12 struct {
	Value []MeasidtoaddmodextR12 `lb:1,ub:maxMeasId`
}
