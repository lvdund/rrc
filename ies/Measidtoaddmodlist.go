package ies

// MeasIdToAddModList ::= SEQUENCE OF MeasIdToAddMod
// SIZE (1..maxNrofMeasId)
type Measidtoaddmodlist struct {
	Value []Measidtoaddmod `lb:1,ub:maxNrofMeasId`
}
