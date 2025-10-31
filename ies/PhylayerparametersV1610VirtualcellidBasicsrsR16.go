package ies

import "rrc/utils"

// PhyLayerParameters-v1610-virtualCellID-BasicSRS-r16 ::= ENUMERATED
type PhylayerparametersV1610VirtualcellidBasicsrsR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610VirtualcellidBasicsrsR16EnumeratedNothing = iota
	PhylayerparametersV1610VirtualcellidBasicsrsR16EnumeratedSupported
)
