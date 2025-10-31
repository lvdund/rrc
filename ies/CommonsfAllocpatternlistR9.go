package ies

// CommonSF-AllocPatternList-r9 ::= SEQUENCE OF MBSFN-SubframeConfig
// SIZE (1..maxMBSFN-Allocations)
type CommonsfAllocpatternlistR9 struct {
	Value []MbsfnSubframeconfig `lb:1,ub:maxMBSFNAllocations`
}
