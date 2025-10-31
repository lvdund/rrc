package ies

// DL-PPW-PreConfigToReleaseList-r17 ::= SEQUENCE OF DL-PPW-ID-r17
// SIZE (1..maxNrofPPW-Config-r17)
type DlPpwPreconfigtoreleaselistR17 struct {
	Value []DlPpwIdR17 `lb:1,ub:maxNrofPPWConfigR17`
}
