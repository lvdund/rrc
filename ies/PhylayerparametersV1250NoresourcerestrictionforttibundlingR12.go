package ies

import "rrc/utils"

// PhyLayerParameters-v1250-noResourceRestrictionForTTIBundling-r12 ::= ENUMERATED
type PhylayerparametersV1250NoresourcerestrictionforttibundlingR12 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1250NoresourcerestrictionforttibundlingR12EnumeratedNothing = iota
	PhylayerparametersV1250NoresourcerestrictionforttibundlingR12EnumeratedSupported
)
