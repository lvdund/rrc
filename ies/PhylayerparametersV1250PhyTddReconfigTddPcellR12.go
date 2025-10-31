package ies

import "rrc/utils"

// PhyLayerParameters-v1250-phy-TDD-ReConfig-TDD-PCell-r12 ::= ENUMERATED
type PhylayerparametersV1250PhyTddReconfigTddPcellR12 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1250PhyTddReconfigTddPcellR12EnumeratedNothing = iota
	PhylayerparametersV1250PhyTddReconfigTddPcellR12EnumeratedSupported
)
