package ies

import "rrc/utils"

// PhyLayerParameters-v1610-virtualCellID-AddSRS-r16 ::= ENUMERATED
type PhylayerparametersV1610VirtualcellidAddsrsR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610VirtualcellidAddsrsR16EnumeratedNothing = iota
	PhylayerparametersV1610VirtualcellidAddsrsR16EnumeratedSupported
)
