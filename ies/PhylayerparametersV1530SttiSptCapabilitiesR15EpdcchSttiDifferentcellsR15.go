package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-epdcch-STTI-differentCells-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15EpdcchSttiDifferentcellsR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15EpdcchSttiDifferentcellsR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15EpdcchSttiDifferentcellsR15EnumeratedSupported
)
