package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-cgi-Acquisition-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16CgiAcquisitionR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16CgiAcquisitionR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16CgiAcquisitionR16EnumeratedSupported
)
