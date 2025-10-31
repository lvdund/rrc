package ies

import "rrc/utils"

// PUSCH-TimeDomainResourceAllocation-mappingType ::= ENUMERATED
type PuschTimedomainresourceallocationMappingtype struct {
	Value utils.ENUMERATED
}

const (
	PuschTimedomainresourceallocationMappingtypeEnumeratedNothing = iota
	PuschTimedomainresourceallocationMappingtypeEnumeratedTypea
	PuschTimedomainresourceallocationMappingtypeEnumeratedTypeb
)
