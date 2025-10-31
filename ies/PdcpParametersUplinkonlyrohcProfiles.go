package ies

import "rrc/utils"

// PDCP-Parameters-uplinkOnlyROHC-Profiles ::= ENUMERATED
type PdcpParametersUplinkonlyrohcProfiles struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersUplinkonlyrohcProfilesEnumeratedNothing = iota
	PdcpParametersUplinkonlyrohcProfilesEnumeratedSupported
)
