package ies

// SPS-ConfigToAddModList-r16 ::= SEQUENCE OF SPS-Config
// SIZE (1..maxNrofSPS-Config-r16)
type SpsConfigtoaddmodlistR16 struct {
	Value []SpsConfig `lb:1,ub:maxNrofSPSConfigR16`
}
