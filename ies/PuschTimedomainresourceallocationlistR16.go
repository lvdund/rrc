package ies

// PUSCH-TimeDomainResourceAllocationList-r16 ::= SEQUENCE OF PUSCH-TimeDomainResourceAllocation-r16
// SIZE (1..maxNrofUL-Allocations-r16)
type PuschTimedomainresourceallocationlistR16 struct {
	Value []PuschTimedomainresourceallocationR16 `lb:1,ub:maxNrofULAllocationsR16`
}
