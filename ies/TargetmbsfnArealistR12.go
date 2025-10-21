package ies

import "rrc/utils"

// TargetMBSFN-AreaList-r12 ::= SEQUENCE OF TargetMBSFN-Area-r12
// SIZE (0..maxMBSFN-Area)
type TargetmbsfnArealistR12 struct {
	Value utils.Sequence[TargetmbsfnAreaR12]
}
