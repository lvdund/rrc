package ies

import "rrc/utils"

// PhyLayerParameters-v1450-ce-SRS-EnhancementWithoutComb4-r14 ::= ENUMERATED
type PhylayerparametersV1450CeSrsEnhancementwithoutcomb4R14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1450CeSrsEnhancementwithoutcomb4R14EnumeratedNothing = iota
	PhylayerparametersV1450CeSrsEnhancementwithoutcomb4R14EnumeratedSupported
)
