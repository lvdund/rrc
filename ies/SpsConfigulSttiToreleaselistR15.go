package ies

// SPS-ConfigUL-STTI-ToReleaseList-r15 ::= SEQUENCE OF SPS-ConfigIndex-r15
// SIZE (1..maxConfigSPS-r15)
type SpsConfigulSttiToreleaselistR15 struct {
	Value []SpsConfigindexR15 `lb:1,ub:maxConfigSPSR15`
}
