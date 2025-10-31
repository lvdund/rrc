package ies

import "rrc/utils"

// Other-Parameters-v1610-resumeWithSCG-Config-r16 ::= ENUMERATED
type OtherParametersV1610ResumewithscgConfigR16 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1610ResumewithscgConfigR16EnumeratedNothing = iota
	OtherParametersV1610ResumewithscgConfigR16EnumeratedSupported
)
