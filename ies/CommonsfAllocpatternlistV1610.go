package ies

// CommonSF-AllocPatternList-v1610 ::= SEQUENCE OF MBSFN-SubframeConfig-v1610
// SIZE (1..maxMBSFN-Allocations)
type CommonsfAllocpatternlistV1610 struct {
	Value []MbsfnSubframeconfigV1610 `lb:1,ub:maxMBSFNAllocations`
}
