package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-dmrs-RepetitionSubslotPDSCH-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15DmrsRepetitionsubslotpdschR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15DmrsRepetitionsubslotpdschR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15DmrsRepetitionsubslotpdschR15EnumeratedSupported
)
