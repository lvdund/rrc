package ies

// MeasResultListMBSFN-r12 ::= SEQUENCE OF MeasResultMBSFN-r12
// SIZE (1..maxMBSFN-Area)
type MeasresultlistmbsfnR12 struct {
	Value []MeasresultmbsfnR12 `lb:1,ub:maxMBSFNArea`
}
