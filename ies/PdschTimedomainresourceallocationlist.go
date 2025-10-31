package ies

// PDSCH-TimeDomainResourceAllocationList ::= SEQUENCE OF PDSCH-TimeDomainResourceAllocation
// SIZE (1..maxNrofDL-Allocations)
type PdschTimedomainresourceallocationlist struct {
	Value []PdschTimedomainresourceallocation `lb:1,ub:maxNrofDLAllocations`
}
