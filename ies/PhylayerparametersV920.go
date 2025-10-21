package ies

import "rrc/utils"

// PhyLayerParameters-v920 ::= SEQUENCE
type PhylayerparametersV920 struct {
	EnhancedduallayerfddR9 *utils.ENUMERATED
	EnhancedduallayertddR9 *utils.ENUMERATED
}
