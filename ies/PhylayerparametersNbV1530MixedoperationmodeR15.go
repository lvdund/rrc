package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1530-mixedOperationMode-r15 ::= ENUMERATED
type PhylayerparametersNbV1530MixedoperationmodeR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1530MixedoperationmodeR15EnumeratedNothing = iota
	PhylayerparametersNbV1530MixedoperationmodeR15EnumeratedSupported
)
