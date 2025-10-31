package ies

// SPS-ConfigSL-ToReleaseList-r14 ::= SEQUENCE OF SPS-ConfigIndex-r14
// SIZE (1..maxConfigSPS-r14)
type SpsConfigslToreleaselistR14 struct {
	Value []SpsConfigindexR14 `lb:1,ub:maxConfigSPSR14`
}
