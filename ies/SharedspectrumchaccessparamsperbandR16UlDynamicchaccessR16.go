package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-ul-DynamicChAccess-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16UlDynamicchaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16UlDynamicchaccessR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16UlDynamicchaccessR16EnumeratedSupported
)
