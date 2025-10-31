package ies

// SPS-ConfigDeactivationStateList-r16 ::= SEQUENCE OF SPS-ConfigDeactivationState-r16
// SIZE (1..maxNrofSPS-DeactivationState)
type SpsConfigdeactivationstatelistR16 struct {
	Value []SpsConfigdeactivationstateR16 `lb:1,ub:maxNrofSPSDeactivationstate`
}
