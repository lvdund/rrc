package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-mib-Acquisition-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16MibAcquisitionR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16MibAcquisitionR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16MibAcquisitionR16EnumeratedSupported
)
