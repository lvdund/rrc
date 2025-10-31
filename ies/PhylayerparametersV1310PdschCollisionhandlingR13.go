package ies

import "rrc/utils"

// PhyLayerParameters-v1310-pdsch-CollisionHandling-r13 ::= ENUMERATED
type PhylayerparametersV1310PdschCollisionhandlingR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1310PdschCollisionhandlingR13EnumeratedNothing = iota
	PhylayerparametersV1310PdschCollisionhandlingR13EnumeratedSupported
)
