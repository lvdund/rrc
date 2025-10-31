package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-ssb-RLM-Semi-StaticChAccess-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16SsbRlmSemiStaticchaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16SsbRlmSemiStaticchaccessR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16SsbRlmSemiStaticchaccessR16EnumeratedSupported
)
