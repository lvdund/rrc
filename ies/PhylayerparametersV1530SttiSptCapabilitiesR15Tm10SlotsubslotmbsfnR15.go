package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-tm10-slotSubslotMBSFN-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15Tm10SlotsubslotmbsfnR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15Tm10SlotsubslotmbsfnR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15Tm10SlotsubslotmbsfnR15EnumeratedSupported
)
