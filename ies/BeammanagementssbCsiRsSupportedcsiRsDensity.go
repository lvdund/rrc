package ies

import "rrc/utils"

// BeamManagementSSB-CSI-RS-supportedCSI-RS-Density ::= ENUMERATED
type BeammanagementssbCsiRsSupportedcsiRsDensity struct {
	Value utils.ENUMERATED
}

const (
	BeammanagementssbCsiRsSupportedcsiRsDensityEnumeratedNothing = iota
	BeammanagementssbCsiRsSupportedcsiRsDensityEnumeratedOne
	BeammanagementssbCsiRsSupportedcsiRsDensityEnumeratedThree
	BeammanagementssbCsiRsSupportedcsiRsDensityEnumeratedOneandthree
)
