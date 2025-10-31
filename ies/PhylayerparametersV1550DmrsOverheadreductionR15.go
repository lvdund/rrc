package ies

import "rrc/utils"

// PhyLayerParameters-v1550-dmrs-OverheadReduction-r15 ::= ENUMERATED
type PhylayerparametersV1550DmrsOverheadreductionR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1550DmrsOverheadreductionR15EnumeratedNothing = iota
	PhylayerparametersV1550DmrsOverheadreductionR15EnumeratedSupported
)
