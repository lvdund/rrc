package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-prach-Wideband-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16PrachWidebandR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16PrachWidebandR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16PrachWidebandR16EnumeratedSupported
)
