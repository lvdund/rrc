package ies

// SPS-ConfigSL-ToAddModList-r14 ::= SEQUENCE OF SPS-ConfigSL-r14
// SIZE (1..maxConfigSPS-r14)
type SpsConfigslToaddmodlistR14 struct {
	Value []SpsConfigslR14 `lb:1,ub:maxConfigSPSR14`
}
