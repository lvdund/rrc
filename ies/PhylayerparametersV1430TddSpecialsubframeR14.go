package ies

import "rrc/utils"

// PhyLayerParameters-v1430-tdd-SpecialSubframe-r14 ::= ENUMERATED
type PhylayerparametersV1430TddSpecialsubframeR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430TddSpecialsubframeR14EnumeratedNothing = iota
	PhylayerparametersV1430TddSpecialsubframeR14EnumeratedSupported
)
