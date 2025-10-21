package ies

import "rrc/utils"

// MBSFN-SubframeConfigList ::= SEQUENCE OF MBSFN-SubframeConfig
// SIZE (1..maxMBSFN-Allocations)
type MbsfnSubframeconfiglist struct {
	Value []MbsfnSubframeconfig
}
