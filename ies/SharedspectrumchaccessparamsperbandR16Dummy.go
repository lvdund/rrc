package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-dummy ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16Dummy struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16DummyEnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16DummyEnumeratedSupported
)
