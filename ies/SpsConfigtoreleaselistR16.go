package ies

// SPS-ConfigToReleaseList-r16 ::= SEQUENCE OF SPS-ConfigIndex-r16
// SIZE (1..maxNrofSPS-Config-r16)
type SpsConfigtoreleaselistR16 struct {
	Value []SpsConfigindexR16 `lb:1,ub:maxNrofSPSConfigR16`
}
