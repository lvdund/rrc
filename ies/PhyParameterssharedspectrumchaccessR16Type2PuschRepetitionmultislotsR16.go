package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-type2-PUSCH-RepetitionMultiSlots-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16Type2PuschRepetitionmultislotsR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16Type2PuschRepetitionmultislotsR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16Type2PuschRepetitionmultislotsR16EnumeratedSupported
)
