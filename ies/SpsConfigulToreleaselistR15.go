package ies

// SPS-ConfigUL-ToReleaseList-r15 ::= SEQUENCE OF SPS-ConfigIndex-r15
// SIZE (1..maxConfigSPS-r15)
type SpsConfigulToreleaselistR15 struct {
	Value []SpsConfigindexR15 `lb:1,ub:maxConfigSPSR15`
}
