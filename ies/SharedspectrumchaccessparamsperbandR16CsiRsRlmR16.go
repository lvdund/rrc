package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-csi-RS-RLM-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16CsiRsRlmR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16CsiRsRlmR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16CsiRsRlmR16EnumeratedSupported
)
