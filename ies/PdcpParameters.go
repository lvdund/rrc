package ies

// PDCP-Parameters ::= SEQUENCE
// Extensible
type PdcpParameters struct {
	SupportedrohcProfiles        RohcProfilesupportlistR15
	MaxnumberrohcContextsessions PdcpParametersMaxnumberrohcContextsessions
}
