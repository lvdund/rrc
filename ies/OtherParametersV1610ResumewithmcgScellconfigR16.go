package ies

import "rrc/utils"

// Other-Parameters-v1610-resumeWithMCG-SCellConfig-r16 ::= ENUMERATED
type OtherParametersV1610ResumewithmcgScellconfigR16 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1610ResumewithmcgScellconfigR16EnumeratedNothing = iota
	OtherParametersV1610ResumewithmcgScellconfigR16EnumeratedSupported
)
