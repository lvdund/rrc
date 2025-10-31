package ies

import "rrc/utils"

// PhyLayerParameters-v1310-pucch-Format5-r13 ::= ENUMERATED
type PhylayerparametersV1310PucchFormat5R13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1310PucchFormat5R13EnumeratedNothing = iota
	PhylayerparametersV1310PucchFormat5R13EnumeratedSupported
)
