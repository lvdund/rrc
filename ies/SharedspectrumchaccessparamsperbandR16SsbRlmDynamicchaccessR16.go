package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-ssb-RLM-DynamicChAccess-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16SsbRlmDynamicchaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16SsbRlmDynamicchaccessR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16SsbRlmDynamicchaccessR16EnumeratedSupported
)
