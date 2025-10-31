package ies

import "rrc/utils"

// PDSCH-TimeDomainResourceAllocation-r16-mappingType-r16 ::= ENUMERATED
type PdschTimedomainresourceallocationR16MappingtypeR16 struct {
	Value utils.ENUMERATED
}

const (
	PdschTimedomainresourceallocationR16MappingtypeR16EnumeratedNothing = iota
	PdschTimedomainresourceallocationR16MappingtypeR16EnumeratedTypea
	PdschTimedomainresourceallocationR16MappingtypeR16EnumeratedTypeb
)
