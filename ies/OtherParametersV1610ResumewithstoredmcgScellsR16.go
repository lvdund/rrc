package ies

import "rrc/utils"

// Other-Parameters-v1610-resumeWithStoredMCG-SCells-r16 ::= ENUMERATED
type OtherParametersV1610ResumewithstoredmcgScellsR16 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1610ResumewithstoredmcgScellsR16EnumeratedNothing = iota
	OtherParametersV1610ResumewithstoredmcgScellsR16EnumeratedSupported
)
