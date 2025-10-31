package ies

// RA-PrioritizationSliceInfo-r17 ::= SEQUENCE
// Extensible
type RaPrioritizationsliceinfoR17 struct {
	NsagIdListR17       []NsagIdR17 `lb:1,ub:maxSliceInfoR17`
	RaPrioritizationR17 RaPrioritization
}
