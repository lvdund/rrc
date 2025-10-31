package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-sib1-Acquisition-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16Sib1AcquisitionR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16Sib1AcquisitionR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16Sib1AcquisitionR16EnumeratedSupported
)
