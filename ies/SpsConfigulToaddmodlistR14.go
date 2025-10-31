package ies

// SPS-ConfigUL-ToAddModList-r14 ::= SEQUENCE OF SPS-ConfigUL
// SIZE (1..maxConfigSPS-r14)
type SpsConfigulToaddmodlistR14 struct {
	Value []SpsConfigul `lb:1,ub:maxConfigSPSR14`
}
