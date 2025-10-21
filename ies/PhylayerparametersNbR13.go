package ies

import "rrc/utils"

// PhyLayerParameters-NB-r13 ::= SEQUENCE
type PhylayerparametersNbR13 struct {
	MultitoneR13    *utils.ENUMERATED
	MulticarrierR13 *utils.ENUMERATED
}
