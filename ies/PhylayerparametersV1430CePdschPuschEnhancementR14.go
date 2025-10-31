package ies

import "rrc/utils"

// PhyLayerParameters-v1430-ce-PDSCH-PUSCH-Enhancement-r14 ::= ENUMERATED
type PhylayerparametersV1430CePdschPuschEnhancementR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430CePdschPuschEnhancementR14EnumeratedNothing = iota
	PhylayerparametersV1430CePdschPuschEnhancementR14EnumeratedSupported
)
