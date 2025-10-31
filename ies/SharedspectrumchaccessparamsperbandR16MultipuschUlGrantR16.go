package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-multiPUSCH-UL-grant-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16MultipuschUlGrantR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16MultipuschUlGrantR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16MultipuschUlGrantR16EnumeratedSupported
)
