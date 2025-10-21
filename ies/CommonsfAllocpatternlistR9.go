package ies

import "rrc/utils"

// CommonSF-AllocPatternList-r9 ::= SEQUENCE OF MBSFN-SubframeConfig
// SIZE (1..maxMBSFN-Allocations)
type CommonsfAllocpatternlistR9 struct {
	Value []MbsfnSubframeconfig
}
