package ies

import "rrc/utils"

// PhyLayerParameters-v1610-widebandPRG-Subframe-r16 ::= ENUMERATED
type PhylayerparametersV1610WidebandprgSubframeR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610WidebandprgSubframeR16EnumeratedNothing = iota
	PhylayerparametersV1610WidebandprgSubframeR16EnumeratedSupported
)
