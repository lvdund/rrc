package ies

import "rrc/utils"

// PhyLayerParameters-v920-enhancedDualLayerTDD-r9 ::= ENUMERATED
type PhylayerparametersV920EnhancedduallayertddR9 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV920EnhancedduallayertddR9EnumeratedNothing = iota
	PhylayerparametersV920EnhancedduallayertddR9EnumeratedSupported
)
