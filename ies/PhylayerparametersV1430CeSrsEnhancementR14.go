package ies

import "rrc/utils"

// PhyLayerParameters-v1430-ce-SRS-Enhancement-r14 ::= ENUMERATED
type PhylayerparametersV1430CeSrsEnhancementR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430CeSrsEnhancementR14EnumeratedNothing = iota
	PhylayerparametersV1430CeSrsEnhancementR14EnumeratedSupported
)
