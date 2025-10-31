package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-configuredGrantWithReTx-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16ConfiguredgrantwithretxR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16ConfiguredgrantwithretxR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16ConfiguredgrantwithretxR16EnumeratedSupported
)
