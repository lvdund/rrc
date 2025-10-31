package ies

import "rrc/utils"

// PhyLayerParameters-v1310-pucch-Format4-r13 ::= ENUMERATED
type PhylayerparametersV1310PucchFormat4R13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1310PucchFormat4R13EnumeratedNothing = iota
	PhylayerparametersV1310PucchFormat4R13EnumeratedSupported
)
