package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-oneShotHARQ-feedback-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16OneshotharqFeedbackR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16OneshotharqFeedbackR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16OneshotharqFeedbackR16EnumeratedSupported
)
