package ies

import "rrc/utils"

// PhyLayerParameters-v1470 ::= SEQUENCE
type PhylayerparametersV1470 struct {
	MimoUeParametersV1470 *MimoUeParametersV1470
	SrsUppts6symR14       *utils.ENUMERATED
}
