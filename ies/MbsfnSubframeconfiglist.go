package ies

// MBSFN-SubframeConfigList ::= SEQUENCE OF MBSFN-SubframeConfig
// SIZE (1..maxMBSFN-Allocations)
type MbsfnSubframeconfiglist struct {
	Value []MbsfnSubframeconfig `lb:1,ub:maxMBSFNAllocations`
}
