package ies

import "rrc/utils"

// Phy-ParametersCommon-pdsch-RepetitionMultiSlots ::= ENUMERATED
type PhyParameterscommonPdschRepetitionmultislots struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPdschRepetitionmultislotsEnumeratedNothing = iota
	PhyParameterscommonPdschRepetitionmultislotsEnumeratedSupported
)
