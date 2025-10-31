package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-pusch-PRB-interlace-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16PuschPrbInterlaceR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16PuschPrbInterlaceR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16PuschPrbInterlaceR16EnumeratedSupported
)
