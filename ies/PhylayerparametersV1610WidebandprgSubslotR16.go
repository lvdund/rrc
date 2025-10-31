package ies

import "rrc/utils"

// PhyLayerParameters-v1610-widebandPRG-Subslot-r16 ::= ENUMERATED
type PhylayerparametersV1610WidebandprgSubslotR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610WidebandprgSubslotR16EnumeratedNothing = iota
	PhylayerparametersV1610WidebandprgSubslotR16EnumeratedSupported
)
