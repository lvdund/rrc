package ies

// CommonSF-AllocPatternList-v1430 ::= SEQUENCE OF MBSFN-SubframeConfig-v1430
// SIZE (1..maxMBSFN-Allocations)
type CommonsfAllocpatternlistV1430 struct {
	Value []MbsfnSubframeconfigV1430 `lb:1,ub:maxMBSFNAllocations`
}
