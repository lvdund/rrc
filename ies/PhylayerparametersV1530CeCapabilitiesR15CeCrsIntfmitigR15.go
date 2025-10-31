package ies

import "rrc/utils"

// PhyLayerParameters-v1530-ce-Capabilities-r15-ce-CRS-IntfMitig-r15 ::= ENUMERATED
type PhylayerparametersV1530CeCapabilitiesR15CeCrsIntfmitigR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530CeCapabilitiesR15CeCrsIntfmitigR15EnumeratedNothing = iota
	PhylayerparametersV1530CeCapabilitiesR15CeCrsIntfmitigR15EnumeratedSupported
)
