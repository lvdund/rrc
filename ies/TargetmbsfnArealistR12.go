package ies

// TargetMBSFN-AreaList-r12 ::= SEQUENCE OF TargetMBSFN-Area-r12
// SIZE (0..maxMBSFN-Area)
type TargetmbsfnArealistR12 struct {
	Value []TargetmbsfnAreaR12 `lb:0,ub:maxMBSFNArea`
}
