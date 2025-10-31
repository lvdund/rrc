package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-ul-AsyncHarqSharingDiff-TTI-Lengths-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15UlAsyncharqsharingdiffTtiLengthsR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15UlAsyncharqsharingdiffTtiLengthsR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15UlAsyncharqsharingdiffTtiLengthsR15EnumeratedSupported
)
