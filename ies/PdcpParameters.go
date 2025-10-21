package ies

import "rrc/utils"

// PDCP-Parameters ::= SEQUENCE
// Extensible
type PdcpParameters struct {
	SupportedrohcProfiles        RohcProfilesupportlistR15
	MaxnumberrohcContextsessions utils.ENUMERATED
}
