package ies

import "rrc/utils"

// Phy-ParametersCommon-type2-PUSCH-RepetitionMultiSlots ::= ENUMERATED
type PhyParameterscommonType2PuschRepetitionmultislots struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonType2PuschRepetitionmultislotsEnumeratedNothing = iota
	PhyParameterscommonType2PuschRepetitionmultislotsEnumeratedSupported
)
