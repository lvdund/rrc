package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-cg-resourceConfig-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16CgResourceconfigR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16CgResourceconfigR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16CgResourceconfigR16EnumeratedSupported
)
