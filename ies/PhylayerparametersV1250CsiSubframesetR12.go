package ies

import "rrc/utils"

// PhyLayerParameters-v1250-csi-SubframeSet-r12 ::= ENUMERATED
type PhylayerparametersV1250CsiSubframesetR12 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1250CsiSubframesetR12EnumeratedNothing = iota
	PhylayerparametersV1250CsiSubframesetR12EnumeratedSupported
)
