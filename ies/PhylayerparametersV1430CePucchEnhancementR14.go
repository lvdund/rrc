package ies

import "rrc/utils"

// PhyLayerParameters-v1430-ce-PUCCH-Enhancement-r14 ::= ENUMERATED
type PhylayerparametersV1430CePucchEnhancementR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430CePucchEnhancementR14EnumeratedNothing = iota
	PhylayerparametersV1430CePucchEnhancementR14EnumeratedSupported
)
