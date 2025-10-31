package ies

// EUTRA-MBSFN-SubframeConfigList ::= SEQUENCE OF EUTRA-MBSFN-SubframeConfig
// SIZE (1..maxMBSFN-Allocations)
type EutraMbsfnSubframeconfiglist struct {
	Value []EutraMbsfnSubframeconfig `lb:1,ub:maxMBSFNAllocations`
}
