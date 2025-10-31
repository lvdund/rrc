package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-powerUCI-SlotPUSCH ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15PoweruciSlotpusch struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15PoweruciSlotpuschEnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15PoweruciSlotpuschEnumeratedSupported
)
