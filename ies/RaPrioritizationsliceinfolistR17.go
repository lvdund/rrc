package ies

// RA-PrioritizationSliceInfoList-r17 ::= SEQUENCE OF RA-PrioritizationSliceInfo-r17
// SIZE (1..maxSliceInfo-r17)
type RaPrioritizationsliceinfolistR17 struct {
	Value []RaPrioritizationsliceinfoR17 `lb:1,ub:maxSliceInfoR17`
}
