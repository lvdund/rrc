package ies

import "rrc/utils"

// PhyLayerParameters-v1610 ::= SEQUENCE
type PhylayerparametersV1610 struct {
	CeCapabilitiesV1610      *interface{}
	WidebandprgSlotR16       *utils.ENUMERATED
	WidebandprgSubslotR16    *utils.ENUMERATED
	WidebandprgSubframeR16   *utils.ENUMERATED
	AddsrsR16                *interface{}
	VirtualcellidBasicsrsR16 *utils.ENUMERATED
	VirtualcellidAddsrsR16   *utils.ENUMERATED
}
