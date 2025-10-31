package ies

// PhyLayerParameters-v1610 ::= SEQUENCE
type PhylayerparametersV1610 struct {
	CeCapabilitiesV1610      *PhylayerparametersV1610CeCapabilitiesV1610
	WidebandprgSlotR16       *PhylayerparametersV1610WidebandprgSlotR16
	WidebandprgSubslotR16    *PhylayerparametersV1610WidebandprgSubslotR16
	WidebandprgSubframeR16   *PhylayerparametersV1610WidebandprgSubframeR16
	AddsrsR16                *PhylayerparametersV1610AddsrsR16
	VirtualcellidBasicsrsR16 *PhylayerparametersV1610VirtualcellidBasicsrsR16
	VirtualcellidAddsrsR16   *PhylayerparametersV1610VirtualcellidAddsrsR16
}
