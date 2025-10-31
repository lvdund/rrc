package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-csi-RS-BFD-CBD-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16CsiRsBfdCbdR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16CsiRsBfdCbdR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16CsiRsBfdCbdR16EnumeratedSupported
)
