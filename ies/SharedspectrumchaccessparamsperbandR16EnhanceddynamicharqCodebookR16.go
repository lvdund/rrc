package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-enhancedDynamicHARQ-codebook-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16EnhanceddynamicharqCodebookR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16EnhanceddynamicharqCodebookR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16EnhanceddynamicharqCodebookR16EnumeratedSupported
)
