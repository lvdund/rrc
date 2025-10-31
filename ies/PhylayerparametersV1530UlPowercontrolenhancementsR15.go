package ies

import "rrc/utils"

// PhyLayerParameters-v1530-ul-PowerControlEnhancements-r15 ::= ENUMERATED
type PhylayerparametersV1530UlPowercontrolenhancementsR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530UlPowercontrolenhancementsR15EnumeratedNothing = iota
	PhylayerparametersV1530UlPowercontrolenhancementsR15EnumeratedSupported
)
