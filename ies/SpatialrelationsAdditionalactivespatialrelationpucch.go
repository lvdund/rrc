package ies

import "rrc/utils"

// SpatialRelations-additionalActiveSpatialRelationPUCCH ::= ENUMERATED
type SpatialrelationsAdditionalactivespatialrelationpucch struct {
	Value utils.ENUMERATED
}

const (
	SpatialrelationsAdditionalactivespatialrelationpucchEnumeratedNothing = iota
	SpatialrelationsAdditionalactivespatialrelationpucchEnumeratedSupported
)
