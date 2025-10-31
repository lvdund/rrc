package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-dmrs-SharingSubslotPDSCH-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15DmrsSharingsubslotpdschR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15DmrsSharingsubslotpdschR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15DmrsSharingsubslotpdschR15EnumeratedSupported
)
