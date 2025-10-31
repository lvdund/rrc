package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-coreset-RB-Offset-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16CoresetRbOffsetR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16CoresetRbOffsetR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16CoresetRbOffsetR16EnumeratedSupported
)
