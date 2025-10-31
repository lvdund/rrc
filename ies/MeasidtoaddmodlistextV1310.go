package ies

// MeasIdToAddModListExt-v1310 ::= SEQUENCE OF MeasIdToAddMod-v1310
// SIZE (1..maxMeasId)
type MeasidtoaddmodlistextV1310 struct {
	Value []MeasidtoaddmodV1310 `lb:1,ub:maxMeasId`
}
