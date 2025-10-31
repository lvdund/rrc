package ies

import "rrc/utils"

// PhyLayerParameters-v1540-cch-IM-RefRecTypeA-OneRX-Port-v1540 ::= ENUMERATED
type PhylayerparametersV1540CchImRefrectypeaOnerxPortV1540 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1540CchImRefrectypeaOnerxPortV1540EnumeratedNothing = iota
	PhylayerparametersV1540CchImRefrectypeaOnerxPortV1540EnumeratedSupported
)
