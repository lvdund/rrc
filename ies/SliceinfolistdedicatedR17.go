package ies

// SliceInfoListDedicated-r17 ::= SEQUENCE OF SliceInfoDedicated-r17
// SIZE (1..maxSliceInfo-r17)
type SliceinfolistdedicatedR17 struct {
	Value []SliceinfodedicatedR17 `lb:1,ub:maxSliceInfoR17`
}
