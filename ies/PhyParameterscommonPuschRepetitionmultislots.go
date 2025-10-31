package ies

import "rrc/utils"

// Phy-ParametersCommon-pusch-RepetitionMultiSlots ::= ENUMERATED
type PhyParameterscommonPuschRepetitionmultislots struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPuschRepetitionmultislotsEnumeratedNothing = iota
	PhyParameterscommonPuschRepetitionmultislotsEnumeratedSupported
)
