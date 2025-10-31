package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-sps-STTI-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15SpsSttiR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15SpsSttiR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15SpsSttiR15EnumeratedSlot
	PhylayerparametersV1530SttiSptCapabilitiesR15SpsSttiR15EnumeratedSubslot
	PhylayerparametersV1530SttiSptCapabilitiesR15SpsSttiR15EnumeratedSlotandsubslot
)
