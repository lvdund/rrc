package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-pucch-F0-F1-PRB-Interlace-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16PucchF0F1PrbInterlaceR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16PucchF0F1PrbInterlaceR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16PucchF0F1PrbInterlaceR16EnumeratedSupported
)
