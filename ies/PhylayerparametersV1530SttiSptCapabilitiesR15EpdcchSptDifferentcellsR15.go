package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-epdcch-SPT-differentCells-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15EpdcchSptDifferentcellsR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15EpdcchSptDifferentcellsR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15EpdcchSptDifferentcellsR15EnumeratedSupported
)
