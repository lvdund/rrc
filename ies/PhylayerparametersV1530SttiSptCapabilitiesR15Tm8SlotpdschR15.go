package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-tm8-slotPDSCH-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15Tm8SlotpdschR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15Tm8SlotpdschR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15Tm8SlotpdschR15EnumeratedSupported
)
