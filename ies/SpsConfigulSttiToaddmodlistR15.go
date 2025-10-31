package ies

// SPS-ConfigUL-STTI-ToAddModList-r15 ::= SEQUENCE OF SPS-ConfigUL-STTI-r15
// SIZE (1..maxConfigSPS-r15)
type SpsConfigulSttiToaddmodlistR15 struct {
	Value []SpsConfigulSttiR15 `lb:1,ub:maxConfigSPSR15`
}
