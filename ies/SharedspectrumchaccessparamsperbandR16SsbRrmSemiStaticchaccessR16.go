package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-ssb-RRM-Semi-StaticChAccess-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16SsbRrmSemiStaticchaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16SsbRrmSemiStaticchaccessR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16SsbRrmSemiStaticchaccessR16EnumeratedSupported
)
