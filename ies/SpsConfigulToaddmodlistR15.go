package ies

// SPS-ConfigUL-ToAddModList-r15 ::= SEQUENCE OF SPS-ConfigUL
// SIZE (1..maxConfigSPS-r15)
type SpsConfigulToaddmodlistR15 struct {
	Value []SpsConfigul `lb:1,ub:maxConfigSPSR15`
}
