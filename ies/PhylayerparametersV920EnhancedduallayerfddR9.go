package ies

import "rrc/utils"

// PhyLayerParameters-v920-enhancedDualLayerFDD-r9 ::= ENUMERATED
type PhylayerparametersV920EnhancedduallayerfddR9 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV920EnhancedduallayerfddR9EnumeratedNothing = iota
	PhylayerparametersV920EnhancedduallayerfddR9EnumeratedSupported
)
