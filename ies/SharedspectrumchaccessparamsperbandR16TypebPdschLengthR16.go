package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-typeB-PDSCH-length-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16TypebPdschLengthR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16TypebPdschLengthR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16TypebPdschLengthR16EnumeratedSupported
)
