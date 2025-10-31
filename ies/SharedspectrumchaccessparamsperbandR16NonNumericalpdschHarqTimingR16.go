package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-non-numericalPDSCH-HARQ-timing-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16NonNumericalpdschHarqTimingR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16NonNumericalpdschHarqTimingR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16NonNumericalpdschHarqTimingR16EnumeratedSupported
)
