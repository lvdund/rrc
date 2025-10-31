package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-tm9-slotSubslot-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15Tm9SlotsubslotR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15Tm9SlotsubslotR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15Tm9SlotsubslotR15EnumeratedSupported
)
