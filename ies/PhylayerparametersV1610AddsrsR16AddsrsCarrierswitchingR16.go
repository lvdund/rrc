package ies

import "rrc/utils"

// PhyLayerParameters-v1610-addSRS-r16-addSRS-CarrierSwitching-r16 ::= ENUMERATED
type PhylayerparametersV1610AddsrsR16AddsrsCarrierswitchingR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610AddsrsR16AddsrsCarrierswitchingR16EnumeratedNothing = iota
	PhylayerparametersV1610AddsrsR16AddsrsCarrierswitchingR16EnumeratedSupported
)
