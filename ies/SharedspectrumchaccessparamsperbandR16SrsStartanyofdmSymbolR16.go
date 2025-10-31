package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-srs-StartAnyOFDM-Symbol-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16SrsStartanyofdmSymbolR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16SrsStartanyofdmSymbolR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16SrsStartanyofdmSymbolR16EnumeratedSupported
)
