package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-type1-PUSCH-RepetitionMultiSlots-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16Type1PuschRepetitionmultislotsR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16Type1PuschRepetitionmultislotsR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16Type1PuschRepetitionmultislotsR16EnumeratedSupported
)
