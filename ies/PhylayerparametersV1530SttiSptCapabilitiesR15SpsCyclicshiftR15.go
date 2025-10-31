package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-sps-cyclicShift-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15SpsCyclicshiftR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15SpsCyclicshiftR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15SpsCyclicshiftR15EnumeratedSupported
)
