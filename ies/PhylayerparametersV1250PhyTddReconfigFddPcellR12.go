package ies

import "rrc/utils"

// PhyLayerParameters-v1250-phy-TDD-ReConfig-FDD-PCell-r12 ::= ENUMERATED
type PhylayerparametersV1250PhyTddReconfigFddPcellR12 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1250PhyTddReconfigFddPcellR12EnumeratedNothing = iota
	PhylayerparametersV1250PhyTddReconfigFddPcellR12EnumeratedSupported
)
