package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-pdsch-RepetitionMultiSlots-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16PdschRepetitionmultislotsR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16PdschRepetitionmultislotsR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16PdschRepetitionmultislotsR16EnumeratedSupported
)
