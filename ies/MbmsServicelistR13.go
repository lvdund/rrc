package ies

// MBMS-ServiceList-r13 ::= SEQUENCE OF MBMS-ServiceInfo-r13
// SIZE (0..maxMBMS-ServiceListPerUE-r13)
type MbmsServicelistR13 struct {
	Value []MbmsServiceinfoR13 `lb:0,ub:maxMBMSServicelistperueR13`
}
