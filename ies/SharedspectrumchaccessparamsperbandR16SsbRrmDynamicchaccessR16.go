package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-ssb-RRM-DynamicChAccess-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16SsbRrmDynamicchaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16SsbRrmDynamicchaccessR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16SsbRrmDynamicchaccessR16EnumeratedSupported
)
