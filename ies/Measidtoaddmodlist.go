package ies

// MeasIdToAddModList ::= SEQUENCE OF MeasIdToAddMod
// SIZE (1..maxMeasId)
type Measidtoaddmodlist struct {
	Value []Measidtoaddmod `lb:1,ub:maxMeasId`
}
