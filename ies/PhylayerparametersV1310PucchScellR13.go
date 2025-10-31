package ies

import "rrc/utils"

// PhyLayerParameters-v1310-pucch-SCell-r13 ::= ENUMERATED
type PhylayerparametersV1310PucchScellR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1310PucchScellR13EnumeratedNothing = iota
	PhylayerparametersV1310PucchScellR13EnumeratedSupported
)
