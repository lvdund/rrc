package ies

import "rrc/utils"

// PDSCH-TimeDomainResourceAllocation-mappingType ::= ENUMERATED
type PdschTimedomainresourceallocationMappingtype struct {
	Value utils.ENUMERATED
}

const (
	PdschTimedomainresourceallocationMappingtypeEnumeratedNothing = iota
	PdschTimedomainresourceallocationMappingtypeEnumeratedTypea
	PdschTimedomainresourceallocationMappingtypeEnumeratedTypeb
)
