package ies

import "rrc/utils"

// PhyLayerParameters ::= SEQUENCE
type Phylayerparameters struct {
	UeTxantennaselectionsupported utils.BOOLEAN
	UeSpecificrefsigssupported    utils.BOOLEAN
}
