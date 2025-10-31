package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-ed-Threshold-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16EdThresholdR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16EdThresholdR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16EdThresholdR16EnumeratedSupported
)
