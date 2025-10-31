package ies

// ConfiguredGrantConfigType2DeactivationStateList-r16 ::= SEQUENCE OF ConfiguredGrantConfigType2DeactivationState-r16
// SIZE (1..maxNrofCG-Type2DeactivationState)
type Configuredgrantconfigtype2deactivationstatelistR16 struct {
	Value []Configuredgrantconfigtype2deactivationstateR16 `lb:1,ub:maxNrofCGType2deactivationstate`
}
