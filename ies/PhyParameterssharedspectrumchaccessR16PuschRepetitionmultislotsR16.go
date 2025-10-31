package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-pusch-RepetitionMultiSlots-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16PuschRepetitionmultislotsR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16PuschRepetitionmultislotsR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16PuschRepetitionmultislotsR16EnumeratedSupported
)
