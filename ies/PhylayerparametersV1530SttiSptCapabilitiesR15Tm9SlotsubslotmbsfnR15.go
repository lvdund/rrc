package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-tm9-slotSubslotMBSFN-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15Tm9SlotsubslotmbsfnR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15Tm9SlotsubslotmbsfnR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15Tm9SlotsubslotmbsfnR15EnumeratedSupported
)
