package ies

import "rrc/utils"

// Phy-ParametersCommon-type1-PUSCH-RepetitionMultiSlots ::= ENUMERATED
type PhyParameterscommonType1PuschRepetitionmultislots struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonType1PuschRepetitionmultislotsEnumeratedNothing = iota
	PhyParameterscommonType1PuschRepetitionmultislotsEnumeratedSupported
)
