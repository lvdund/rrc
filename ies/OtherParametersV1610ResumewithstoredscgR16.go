package ies

import "rrc/utils"

// Other-Parameters-v1610-resumeWithStoredSCG-r16 ::= ENUMERATED
type OtherParametersV1610ResumewithstoredscgR16 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1610ResumewithstoredscgR16EnumeratedNothing = iota
	OtherParametersV1610ResumewithstoredscgR16EnumeratedSupported
)
