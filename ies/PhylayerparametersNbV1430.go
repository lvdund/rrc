package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1430 ::= SEQUENCE
type PhylayerparametersNbV1430 struct {
	MulticarrierNprachR14 *utils.ENUMERATED
	TwoharqProcessesR14   *utils.ENUMERATED
}
