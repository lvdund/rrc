package ies

import "rrc/utils"

// PhyLayerParameters-v1130-tdd-SpecialSubframe-r11 ::= ENUMERATED
type PhylayerparametersV1130TddSpecialsubframeR11 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1130TddSpecialsubframeR11EnumeratedNothing = iota
	PhylayerparametersV1130TddSpecialsubframeR11EnumeratedSupported
)
