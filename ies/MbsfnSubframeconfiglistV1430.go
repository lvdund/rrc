package ies

// MBSFN-SubframeConfigList-v1430 ::= SEQUENCE OF MBSFN-SubframeConfig-v1430
// SIZE (1..maxMBSFN-Allocations)
type MbsfnSubframeconfiglistV1430 struct {
	Value []MbsfnSubframeconfigV1430 `lb:1,ub:maxMBSFNAllocations`
}
