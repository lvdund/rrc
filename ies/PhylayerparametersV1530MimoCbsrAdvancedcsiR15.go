package ies

import "rrc/utils"

// PhyLayerParameters-v1530-mimo-CBSR-AdvancedCSI-r15 ::= ENUMERATED
type PhylayerparametersV1530MimoCbsrAdvancedcsiR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530MimoCbsrAdvancedcsiR15EnumeratedNothing = iota
	PhylayerparametersV1530MimoCbsrAdvancedcsiR15EnumeratedSupported
)
