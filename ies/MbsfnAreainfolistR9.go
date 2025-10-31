package ies

// MBSFN-AreaInfoList-r9 ::= SEQUENCE OF MBSFN-AreaInfo-r9
// SIZE (1..maxMBSFN-Area)
type MbsfnAreainfolistR9 struct {
	Value []MbsfnAreainfoR9 `lb:1,ub:maxMBSFNArea`
}
