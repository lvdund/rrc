package ies

// PDSCH-TimeDomainResourceAllocationList-r16 ::= SEQUENCE OF PDSCH-TimeDomainResourceAllocation-r16
// SIZE (1..maxNrofDL-Allocations)
type PdschTimedomainresourceallocationlistR16 struct {
	Value []PdschTimedomainresourceallocationR16 `lb:1,ub:maxNrofDLAllocations`
}
