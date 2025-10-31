package ies

import "rrc/utils"

// PhyLayerParameters-v1610-widebandPRG-Slot-r16 ::= ENUMERATED
type PhylayerparametersV1610WidebandprgSlotR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610WidebandprgSlotR16EnumeratedNothing = iota
	PhylayerparametersV1610WidebandprgSlotR16EnumeratedSupported
)
