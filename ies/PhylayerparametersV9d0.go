package ies

import "rrc/utils"

// PhyLayerParameters-v9d0 ::= SEQUENCE
type PhylayerparametersV9d0 struct {
	Tm5FddR9 *utils.ENUMERATED
	Tm5TddR9 *utils.ENUMERATED
}
