package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-ul-Semi-StaticChAccess-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16UlSemiStaticchaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16UlSemiStaticchaccessR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16UlSemiStaticchaccessR16EnumeratedSupported
)
