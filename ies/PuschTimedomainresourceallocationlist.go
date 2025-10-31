package ies

// PUSCH-TimeDomainResourceAllocationList ::= SEQUENCE OF PUSCH-TimeDomainResourceAllocation
// SIZE (1..maxNrofUL-Allocations)
type PuschTimedomainresourceallocationlist struct {
	Value []PuschTimedomainresourceallocation `lb:1,ub:maxNrofULAllocations`
}
