package ies

// SPS-ConfigUL-ToReleaseList-r14 ::= SEQUENCE OF SPS-ConfigIndex-r14
// SIZE (1..maxConfigSPS-r14)
type SpsConfigulToreleaselistR14 struct {
	Value []SpsConfigindexR14 `lb:1,ub:maxConfigSPSR14`
}
