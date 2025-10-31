package ies

import "rrc/utils"

// PhyLayerParameters-v1530-ce-Capabilities-r15-ce-CQI-AlternativeTable-r15 ::= ENUMERATED
type PhylayerparametersV1530CeCapabilitiesR15CeCqiAlternativetableR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530CeCapabilitiesR15CeCqiAlternativetableR15EnumeratedNothing = iota
	PhylayerparametersV1530CeCapabilitiesR15CeCqiAlternativetableR15EnumeratedSupported
)
