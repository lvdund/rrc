package ies

// SPS-Config-v1430 ::= SEQUENCE
type SpsConfigV1430 struct {
	UlSpsVRntiR14               *CRnti
	SlSpsVRntiR14               *CRnti
	SpsConfigulToaddmodlistR14  *SpsConfigulToaddmodlistR14
	SpsConfigulToreleaselistR14 *SpsConfigulToreleaselistR14
	SpsConfigslToaddmodlistR14  *SpsConfigslToaddmodlistR14
	SpsConfigslToreleaselistR14 *SpsConfigslToreleaselistR14
}
