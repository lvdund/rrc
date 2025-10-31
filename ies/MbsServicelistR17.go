package ies

// MBS-ServiceList-r17 ::= SEQUENCE OF MBS-ServiceInfo-r17
// SIZE (1..maxNrofMBS-ServiceListPerUE-r17)
type MbsServicelistR17 struct {
	Value []MbsServiceinfoR17 `lb:1,ub:maxNrofMBSServicelistperueR17`
}
