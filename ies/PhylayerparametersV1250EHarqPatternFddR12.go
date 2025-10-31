package ies

import "rrc/utils"

// PhyLayerParameters-v1250-e-HARQ-Pattern-FDD-r12 ::= ENUMERATED
type PhylayerparametersV1250EHarqPatternFddR12 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1250EHarqPatternFddR12EnumeratedNothing = iota
	PhylayerparametersV1250EHarqPatternFddR12EnumeratedSupported
)
