package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-powerUCI-SubslotPUSCH ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15PoweruciSubslotpusch struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15PoweruciSubslotpuschEnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15PoweruciSubslotpuschEnumeratedSupported
)
