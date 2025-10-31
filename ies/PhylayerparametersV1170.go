package ies

import "rrc/utils"

// PhyLayerParameters-v1170 ::= SEQUENCE
type PhylayerparametersV1170 struct {
	InterbandtddCaWithdifferentconfigR11 *utils.BITSTRING `lb:2,ub:2`
}
